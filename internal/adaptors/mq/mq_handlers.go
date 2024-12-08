package mq

import (
	"chat-backend-general/internal/domain"
	usecases "chat-backend-general/internal/usecases/mq"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type MessageQueueHandler struct {
	useCase usecases.MessageQueueUseCase // Use interface directly
}

// NewMessageQueueHandler creates a new handler with the provided use case
func NewMessageQueueHandler(useCase usecases.MessageQueueUseCase) *MessageQueueHandler {
	return &MessageQueueHandler{useCase: useCase}
}

// PublishMessage handles the publishing of messages to the message queue
func (h *MessageQueueHandler) PublishMessage(c *gin.Context) {
	// Define the request structure
	var request struct {
		Task   string                 `json:"task" binding:"required"`
		Args   []interface{}          `json:"args"`
		Kwargs map[string]interface{} `json:"kwargs"`
		ETA    *string                `json:"eta,omitempty"`
	}

	// Bind and validate the JSON request payload
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request payload",
			"details": err.Error(),
		})
		return
	}

	// Extract queue name or use default
	queueName := c.DefaultQuery("queueName", "default")

	// Construct Celery-compatible message
	message := domain.NewCeleryMessage(request.Task, request.Args, request.Kwargs)
	if request.ETA != nil {
		// Optionally set the ETA
		parsedETA, err := parseETA(*request.ETA)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid ETA format",
				"details": err.Error(),
			})
			return
		}
		message.ETA = &parsedETA
	}

	// Publish the message
	if err := h.useCase.Publish(queueName, message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to publish message",
			"details": err.Error(),
		})
		return
	}

	// Success response
	c.JSON(http.StatusOK, gin.H{
		"status":    "Message published successfully",
		"messageID": message.ID,
	})
}

// parseETA parses the ETA string into a time.Time object
func parseETA(eta string) (time.Time, error) {
	parsedETA, err := time.Parse(time.RFC3339, eta)
	if err != nil {
		return time.Time{}, err
	}
	return parsedETA, nil
}
