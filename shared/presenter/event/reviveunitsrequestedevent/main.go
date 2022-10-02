package reviveunitsrequestedevent

import (
	"encoding/json"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/google/uuid"
)

type localCoordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type payload struct {
	GameId      uuid.UUID         `json:"gameId"`
	Coordinates []localCoordinate `json:"coordinates"`
}

type Event struct {
	Payload   payload   `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
}

func NewEventTopic() string {
	return "game-room-revive-units-requested"
}

func NewEvent(gameId uuid.UUID, coordinates []valueobject.Coordinate) []byte {
	payloadCoordinates := make([]localCoordinate, 0)

	for _, coordinate := range coordinates {
		payloadCoordinates = append(payloadCoordinates, localCoordinate{X: coordinate.GetX(), Y: coordinate.GetY()})
	}

	event := Event{
		Payload: payload{
			GameId:      gameId,
			Coordinates: payloadCoordinates,
		},
		Timestamp: time.Now(),
	}

	jsonBytes, _ := json.Marshal(event)

	return jsonBytes
}
