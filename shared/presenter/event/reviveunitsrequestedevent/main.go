package reviveunitsrequestedevent

import (
	"encoding/json"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/dto/coordinatedto"
	"github.com/google/uuid"
)

type payload struct {
	GameId      uuid.UUID           `json:"gameId"`
	Coordinates []coordinatedto.Dto `json:"coordinates"`
}

type Event struct {
	Payload   payload   `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
}

func NewEventTopic() string {
	return "game-room-revive-units-requested"
}

func NewEvent(gameId uuid.UUID, coordinates []valueobject.Coordinate) []byte {
	coordinateDtos := coordinatedto.ToDtoList(coordinates)

	event := Event{
		Payload: payload{
			GameId:      gameId,
			Coordinates: coordinateDtos,
		},
		Timestamp: time.Now(),
	}

	jsonBytes, _ := json.Marshal(event)

	return jsonBytes
}
