package gameinfoupdatedevent

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/dto/mapsizedto"
	"github.com/google/uuid"
)

type payload struct {
	MapSize mapsizedto.Dto `json:"mapSize"`
}

type Event struct {
	Payload   payload   `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
}

func NewEventTopic(gameId uuid.UUID, playerId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-player-%s-game-info-updated", gameId, playerId)
}

func NewEvent(mapSize valueobject.MapSize) []byte {
	mapSizeDto := mapsizedto.ToDto(mapSize)

	event := Event{
		Payload: payload{
			MapSize: mapSizeDto,
		},
		Timestamp: time.Now(),
	}

	jsonBytes, _ := json.Marshal(event)

	return jsonBytes
}
