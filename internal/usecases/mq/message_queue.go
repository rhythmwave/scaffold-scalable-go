package mq

import "chat-backend-general/internal/domain"

// MessageQueueUseCase defines the interface for message queue use cases
type MessageQueueUseCase interface {
    Publish(queueName string, payload domain.CeleryMessage) error
}

// messageQueueUseCaseImpl is the concrete implementation of MessageQueueUseCase
type messageQueueUseCaseImpl struct {
    queue domain.MessageQueue
}

// NewMessageQueueUseCase creates a new instance of MessageQueueUseCase
func NewMessageQueueUseCase(queue domain.MessageQueue) MessageQueueUseCase {
    return &messageQueueUseCaseImpl{queue: queue}
}

// Publish sends a Celery-compatible message to the specified queue
func (m *messageQueueUseCaseImpl) Publish(queueName string, payload domain.CeleryMessage) error {
    return m.queue.PublishMessage(queueName, payload)
}
