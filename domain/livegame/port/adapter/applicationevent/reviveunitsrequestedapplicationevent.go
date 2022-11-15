package applicationevent

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/presenter/presenterdto"
	"github.com/google/uuid"
)

type ReviveUnitsRequestedApplicationEvent struct {
	Coordinates []presenterdto.CoordinatePresenterDto `json:"coordinates"`
}

func NewReviveUnitsRequestedApplicationEventTopic(liveGameId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-revive-units-requested", liveGameId)
}

func NewReviveUnitsRequestedApplicationEvent(coordinatePresenterDtos []presenterdto.CoordinatePresenterDto) ReviveUnitsRequestedApplicationEvent {
	return ReviveUnitsRequestedApplicationEvent{
		Coordinates: coordinatePresenterDtos,
	}
}
