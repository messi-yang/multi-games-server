package unitmaptickedevent

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type payload struct{}

type Event struct {
	Payload   payload   `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
}

func NewEventTopic(gameId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-unit-revived", gameId)
}

func NewEvent() []byte {
	event := Event{
		Payload:   payload{},
		Timestamp: time.Now(),
	}

	jsonBytes, _ := json.Marshal(event)

	return jsonBytes
}
