package applicationevent

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/domain/livegame/port/dto"
	"github.com/google/uuid"
)

type ReviveUnitsRequestedApplicationEvent struct {
	Coordinates []dto.CoordinateDto `json:"coordinates"`
}

func NewReviveUnitsRequestedApplicationEventTopic(liveGameId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-revive-units-requested", liveGameId)
}

func NewReviveUnitsRequestedApplicationEvent(coordinateDtos []dto.CoordinateDto) ReviveUnitsRequestedApplicationEvent {
	return ReviveUnitsRequestedApplicationEvent{
		Coordinates: coordinateDtos,
	}
}
