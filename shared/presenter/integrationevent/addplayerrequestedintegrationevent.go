package integrationevent

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type addPlayerRequestedIntegrationEventPayload struct {
	PlayerId uuid.UUID `json:"playerId"`
}

type AddPlayerRequestedIntegrationEvent struct {
	Payload   addPlayerRequestedIntegrationEventPayload `json:"payload"`
	Timestamp time.Time                                 `json:"timestamp"`
}

func NewAddPlayerRequestedIntegrationEventTopic(gameId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-add-player-requested", gameId)
}

func NewAddPlayerRequestedIntegrationEvent(playerId uuid.UUID) []byte {
	event := AddPlayerRequestedIntegrationEvent{
		Payload: addPlayerRequestedIntegrationEventPayload{
			PlayerId: playerId,
		},
		Timestamp: time.Now(),
	}

	jsonBytes, _ := json.Marshal(event)

	return jsonBytes
}
