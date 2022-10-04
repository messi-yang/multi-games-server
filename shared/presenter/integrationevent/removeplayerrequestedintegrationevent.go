package integrationevent

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type removePlayerRequestedIntegrationEventPayload struct {
	PlayerId uuid.UUID `json:"playerId"`
}

type RemovePlayerRequestedIntegrationEvent struct {
	Payload   removePlayerRequestedIntegrationEventPayload `json:"payload"`
	Timestamp time.Time                                    `json:"timestamp"`
}

func NewRemovePlayerRequestedIntegrationEventTopic(gameId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-remove-player-requested", gameId)
}

func NewRemovePlayerRequestedIntegrationEvent(playerId uuid.UUID) []byte {
	event := RemovePlayerRequestedIntegrationEvent{
		Payload: removePlayerRequestedIntegrationEventPayload{
			PlayerId: playerId,
		},
		Timestamp: time.Now(),
	}

	jsonBytes, _ := json.Marshal(event)

	return jsonBytes
}
