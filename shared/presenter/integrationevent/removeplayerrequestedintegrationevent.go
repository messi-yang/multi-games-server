package integrationevent

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game/dto"
	"github.com/google/uuid"
)

type removePlayerRequestedIntegrationEventPayload struct {
	PlayerId dto.PlayerIdDto `json:"playerId"`
}

type RemovePlayerRequestedIntegrationEvent struct {
	Payload   removePlayerRequestedIntegrationEventPayload `json:"payload"`
	Timestamp time.Time                                    `json:"timestamp"`
}

func NewRemovePlayerRequestedIntegrationEventTopic(gameId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-remove-player-requested", gameId)
}

func NewRemovePlayerRequestedIntegrationEvent(playerIdDto dto.PlayerIdDto) []byte {
	event := RemovePlayerRequestedIntegrationEvent{
		Payload: removePlayerRequestedIntegrationEventPayload{
			PlayerId: playerIdDto,
		},
		Timestamp: time.Now(),
	}

	jsonBytes, _ := json.Marshal(event)

	return jsonBytes
}
