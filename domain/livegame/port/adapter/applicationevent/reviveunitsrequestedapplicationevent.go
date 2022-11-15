package applicationevent

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/uipresenter/uidto"
	"github.com/google/uuid"
)

type ReviveUnitsRequestedApplicationEvent struct {
	Coordinates []uidto.CoordinateUiDto `json:"coordinates"`
}

func NewReviveUnitsRequestedApplicationEventTopic(liveGameId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-revive-units-requested", liveGameId)
}

func NewReviveUnitsRequestedApplicationEvent(coordinateUiDtos []uidto.CoordinateUiDto) ReviveUnitsRequestedApplicationEvent {
	return ReviveUnitsRequestedApplicationEvent{
		Coordinates: coordinateUiDtos,
	}
}
