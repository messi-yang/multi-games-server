package removeplayerrequestedevent

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type payload struct {
	PlayerId uuid.UUID `json:"playerId"`
}

type Event struct {
	Payload   payload   `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
}

func NewEventTopic(gameId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-remove-player-requested", gameId)
}

func NewEvent(playerId uuid.UUID) []byte {
	event := Event{
		Payload: payload{
			PlayerId: playerId,
		},
		Timestamp: time.Now(),
	}

	jsonBytes, _ := json.Marshal(event)

	return jsonBytes
}
