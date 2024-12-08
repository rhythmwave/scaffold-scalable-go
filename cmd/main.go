package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"chat-backend-general/config"
	httpInfra "chat-backend-general/internal/infra/http"

	"go.uber.org/zap"
)

func main() {
	// Initialize logger
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync() // Flushes buffer, if any

	// Initialize config
	cfg, err := config.Init()
	if err != nil {
		logger.Fatal("Failed to initialize config", zap.Error(err))
	}

	// Create and run Gin server
	server := httpInfra.NewGinServer(cfg, logger)
	go func() {
		logger.Info("Starting server on :8080")
		if err := server.Run(":8080"); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Create HTTP server instance
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: server,
	}

	<-quit
	logger.Info("Shutting down server...")

	// Create a context with timeout to allow current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		logger.Fatal("Failed to gracefully stop server", zap.Error(err))
	}

	logger.Info("Server stopped")
}
