package unitsrevivedevent

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/dto/coordinatedto"
	"github.com/google/uuid"
)

type payload struct {
	Coordinates []coordinatedto.Dto `json:"coordinates"`
}

type Event struct {
	Payload   payload   `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
}

func NewEventTopic(gameId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-units-revived", gameId)
}

func NewEvent(coordinates []valueobject.Coordinate) []byte {
	coordinateDtos := coordinatedto.ToDtoList(coordinates)

	event := Event{
		Payload: payload{
			Coordinates: coordinateDtos,
		},
		Timestamp: time.Now(),
	}

	jsonBytes, _ := json.Marshal(event)

	return jsonBytes
}
