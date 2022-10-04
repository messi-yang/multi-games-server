package reviveunitsrequestedevent

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/dto"
	"github.com/google/uuid"
)

type payload struct {
	Coordinates []dto.CoordinateDto `json:"coordinates"`
}

type Event struct {
	Payload   payload   `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
}

func NewEventTopic(gameId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-revive-units-requested", gameId)
}

func NewEvent(coordinates []valueobject.Coordinate) []byte {
	coordinateDtos := dto.NewCoordinateDtos(coordinates)

	event := Event{
		Payload: payload{
			Coordinates: coordinateDtos,
		},
		Timestamp: time.Now(),
	}

	jsonBytes, _ := json.Marshal(event)

	return jsonBytes
}
