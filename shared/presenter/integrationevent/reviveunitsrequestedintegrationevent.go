package integrationevent

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/dto"
	"github.com/google/uuid"
)

type reviveUnitsRequestedIntegrationEventPayload struct {
	Coordinates []dto.CoordinateDto `json:"coordinates"`
}

type ReviveUnitsRequestedIntegrationEvent struct {
	Payload   reviveUnitsRequestedIntegrationEventPayload `json:"payload"`
	Timestamp time.Time                                   `json:"timestamp"`
}

func NewReviveUnitsRequestedIntegrationEventTopic(gameId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-revive-units-requested", gameId)
}

func NewReviveUnitsRequestedIntegrationEvent(coordinates []valueobject.Coordinate) []byte {
	coordinateDtos := dto.NewCoordinateDtos(coordinates)

	event := ReviveUnitsRequestedIntegrationEvent{
		Payload: reviveUnitsRequestedIntegrationEventPayload{
			Coordinates: coordinateDtos,
		},
		Timestamp: time.Now(),
	}

	jsonBytes, _ := json.Marshal(event)

	return jsonBytes
}
