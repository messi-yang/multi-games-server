package integrationevent

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game/dto"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game/valueobject"
	"github.com/google/uuid"
)

type gameInfoUpdatedIntegrationEventPayload struct {
	MapSize dto.MapSizeDto `json:"mapSize"`
}

type GameInfoUpdatedIntegrationEvent struct {
	Payload   gameInfoUpdatedIntegrationEventPayload `json:"payload"`
	Timestamp time.Time                              `json:"timestamp"`
}

func NewGameInfoUpdatedIntegrationEventTopic(gameId uuid.UUID, playerId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-player-%s-game-info-updated", gameId, playerId)
}

func NewGameInfoUpdatedIntegrationEvent(mapSize valueobject.MapSize) []byte {
	mapSizeDto := dto.NewMapSizeDto(mapSize)

	event := GameInfoUpdatedIntegrationEvent{
		Payload: gameInfoUpdatedIntegrationEventPayload{
			MapSize: mapSizeDto,
		},
		Timestamp: time.Now(),
	}

	jsonBytes, _ := json.Marshal(event)

	return jsonBytes
}
