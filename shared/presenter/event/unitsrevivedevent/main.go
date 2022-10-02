package unitsrevivedevent

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/google/uuid"
)

type localCoordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type payload struct {
	Coordinates []localCoordinate `json:"coordinates"`
}

type Event struct {
	Payload   payload   `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
}

func NewEventTopic(gameId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-units-revived", gameId)
}

func NewEvent(coordinates []valueobject.Coordinate) []byte {
	payloadCoordinates := make([]localCoordinate, 0)

	for _, coordinate := range coordinates {
		payloadCoordinates = append(payloadCoordinates, localCoordinate{X: coordinate.GetX(), Y: coordinate.GetY()})
	}

	event := Event{
		Payload: payload{
			Coordinates: payloadCoordinates,
		},
		Timestamp: time.Now(),
	}

	jsonBytes, _ := json.Marshal(event)

	return jsonBytes
}
