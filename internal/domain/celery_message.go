package domain

import (
	"time"

	"github.com/google/uuid"
)

type CeleryMessage struct {
	Task    string                 `json:"task"`
	Args    []interface{}          `json:"args"`
	Kwargs  map[string]interface{} `json:"kwargs"`
	ID      string                 `json:"id"`
	ETA     *time.Time             `json:"eta,omitempty"`
	Expires *time.Time             `json:"expires,omitempty"`
}

// NewCeleryMessage creates a Celery-compatible message
func NewCeleryMessage(task string, args []interface{}, kwargs map[string]interface{}) CeleryMessage {
	return CeleryMessage{
		Task:   task,
		Args:   args,
		Kwargs: kwargs,
		ID:     uuid.New().String(), // Generate a unique ID
	}
}
