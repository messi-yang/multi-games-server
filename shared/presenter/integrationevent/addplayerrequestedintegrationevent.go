package integrationevent

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game/dto"
	"github.com/google/uuid"
)

type addPlayerRequestedIntegrationEventPayload struct {
	PlayerId dto.PlayerIdDto `json:"playerId"`
}

type AddPlayerRequestedIntegrationEvent struct {
	Payload   addPlayerRequestedIntegrationEventPayload `json:"payload"`
	Timestamp time.Time                                 `json:"timestamp"`
}

func NewAddPlayerRequestedIntegrationEventTopic(gameId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-add-player-requested", gameId)
}

func NewAddPlayerRequestedIntegrationEvent(playerIdDto dto.PlayerIdDto) []byte {
	event := AddPlayerRequestedIntegrationEvent{
		Payload: addPlayerRequestedIntegrationEventPayload{
			PlayerId: playerIdDto,
		},
		Timestamp: time.Now(),
	}

	jsonBytes, _ := json.Marshal(event)

	return jsonBytes
}
