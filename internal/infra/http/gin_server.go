package http

import (
	"chat-backend-general/config"
	usecasesHttp "chat-backend-general/internal/adaptors/http"
	usecasesMq "chat-backend-general/internal/adaptors/mq"
	usecasesStorage "chat-backend-general/internal/adaptors/storage"
	usecasesValidation "chat-backend-general/internal/adaptors/validation"
	"chat-backend-general/internal/domain"
	usecasesFileUpload "chat-backend-general/internal/usecases"
	usecasesMqConcrete "chat-backend-general/internal/usecases/mq"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewGinServer(cfg *config.Config, logger *zap.Logger) *gin.Engine {
	r := gin.Default()

	// Middleware
	r.Use(cors.Default())

	// Initialize storage adapter and file upload use case
	storageAdapter, err := usecasesStorage.NewBlobStorageAdapter(cfg, logger)
	if err != nil {
		logger.Fatal("Error initializing storage adapter", zap.Error(err))
	}
	fileUploadUseCase := usecasesFileUpload.NewFileUploadUseCase(storageAdapter)

	// Initialize file validators
	fileSizeValidator := usecasesValidation.NewFileSizeValidator(10 * 1024 * 1024) // 10MB
	fileTypeValidator := usecasesValidation.NewFileTypeValidator([]string{
		"application/msword", "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		"application/vnd.ms-powerpoint", "application/vnd.openxmlformats-officedocument.presentationml.presentation",
		"application/vnd.ms-excel", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		"application/vnd.oasis.opendocument.text", "application/vnd.oasis.opendocument.presentation", "application/vnd.oasis.opendocument.spreadsheet",
		"text/plain", "application/pdf",
		"text/csv", "text/html",
		"application/xliff+xml",
		"text/markdown",
		"application/vnd.ms-outlook",
		"text/rtf",
		"text/tab-separated-values", "text/tab-separated-values",
	})

	fileValidators := []domain.FileValidator{
		fileSizeValidator,
		fileTypeValidator,
	}

	fileHandler := usecasesHttp.NewFileHandler(fileUploadUseCase, fileValidators)

	// Initialize Azure Service Bus adapter
	messageQueueAdapter, err := usecasesMq.NewAzureServiceBusAdapter(cfg.ServiceBus.ConnectionString, logger)
	if err != nil {
		logger.Fatal("Failed to initialize Azure Service Bus adapter", zap.Error(err))
	}
	messageQueueUseCase := usecasesMqConcrete.NewMessageQueueUseCase(messageQueueAdapter)
	messageQueueHandler := usecasesMq.NewMessageQueueHandler(messageQueueUseCase)

	// Define file upload endpoint
	r.POST("/doc/upload", fileHandler.UploadFile)

	// Define message queue endpoints
	r.POST("/queue/publish", messageQueueHandler.PublishMessage)

	return r
}
