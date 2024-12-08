package domain

type MessageQueue interface {
	PublishMessage(queueName string, message CeleryMessage) error
}
