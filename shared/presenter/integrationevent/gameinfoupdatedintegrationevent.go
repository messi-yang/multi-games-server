package integrationevent

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game/dto"
	"github.com/google/uuid"
)

type gameInfoUpdatedIntegrationEventPayload struct {
	Dimension dto.DimensionDto `json:"dimension"`
}

type GameInfoUpdatedIntegrationEvent struct {
	Payload   gameInfoUpdatedIntegrationEventPayload `json:"payload"`
	Timestamp time.Time                              `json:"timestamp"`
}

func NewGameInfoUpdatedIntegrationEventTopic(gameId uuid.UUID, playerIdDto dto.PlayerIdDto) string {
	return fmt.Sprintf("game-room-%s-player-%s-game-info-updated", gameId, playerIdDto.Value)
}

func NewGameInfoUpdatedIntegrationEvent(dimensionDto dto.DimensionDto) []byte {
	event := GameInfoUpdatedIntegrationEvent{
		Payload: gameInfoUpdatedIntegrationEventPayload{
			Dimension: dimensionDto,
		},
		Timestamp: time.Now(),
	}

	jsonBytes, _ := json.Marshal(event)

	return jsonBytes
}
