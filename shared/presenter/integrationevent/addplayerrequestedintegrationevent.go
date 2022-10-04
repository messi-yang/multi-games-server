package integrationevent

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/entity"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/dto"
	"github.com/google/uuid"
)

type addPlayerRequestedIntegrationEventPayload struct {
	Player dto.PlayerDto `json:"player"`
}

type AddPlayerRequestedIntegrationEvent struct {
	Payload   addPlayerRequestedIntegrationEventPayload `json:"payload"`
	Timestamp time.Time                                 `json:"timestamp"`
}

func NewAddPlayerRequestedIntegrationEventTopic(gameId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-add-player-requested", gameId)
}

func NewAddPlayerRequestedIntegrationEvent(player entity.Player) []byte {
	playerDto := dto.NewPlayerDto(player)

	event := AddPlayerRequestedIntegrationEvent{
		Payload: addPlayerRequestedIntegrationEventPayload{
			Player: playerDto,
		},
		Timestamp: time.Now(),
	}

	jsonBytes, _ := json.Marshal(event)

	return jsonBytes
}
