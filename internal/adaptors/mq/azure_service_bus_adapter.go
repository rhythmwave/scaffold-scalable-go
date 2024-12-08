package mq

import (
	"chat-backend-general/internal/domain"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"go.uber.org/zap"
)

// AzureServiceBusAdapter is an adapter for Azure Service Bus.
type AzureServiceBusAdapter struct {
	client *azservicebus.Client
	logger *zap.Logger
}

// NewAzureServiceBusAdapter initializes a new AzureServiceBusAdapter.
func NewAzureServiceBusAdapter(connectionString string, logger *zap.Logger) (*AzureServiceBusAdapter, error) {
	client, err := azservicebus.NewClientFromConnectionString(connectionString, nil)
	if err != nil {
		logger.Error("Failed to create service bus client", zap.Error(err))
		return nil, fmt.Errorf("failed to create service bus client: %w", err)
	}

	return &AzureServiceBusAdapter{
		client: client,
		logger: logger,
	}, nil
}

// PublishMessage publishes a message to the specified Azure Service Bus queue.
func (a *AzureServiceBusAdapter) PublishMessage(queueName string, message domain.CeleryMessage) error {
	if a.client == nil {
		a.logger.Error("Service Bus client is nil")
		return errors.New("service bus client is nil")
	}

	sender, err := a.client.NewSender(queueName, nil)
	if err != nil {
		a.logger.Error("Failed to create sender", zap.Error(err), zap.String("queueName", queueName))
		return fmt.Errorf("failed to create sender: %w", err)
	}
	defer sender.Close(context.Background())

	messageBytes, err := json.Marshal(message)
	if err != nil {
		a.logger.Error("Failed to marshal message", zap.Error(err))
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	ttl := time.Hour * 1

	sbMessage := &azservicebus.Message{
		Body:       messageBytes,
		TimeToLive: &ttl, // Message TTL
	}

	if err := sender.SendMessage(context.Background(), sbMessage, nil); err != nil {
		a.logger.Error("Failed to send message", zap.Error(err), zap.String("queueName", queueName))
		return fmt.Errorf("failed to send message: %w", err)
	}

	a.logger.Info("Message sent successfully", zap.String("queueName", queueName))
	return nil
}
