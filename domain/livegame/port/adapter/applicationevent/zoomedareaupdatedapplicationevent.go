package applicationevent

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/presenter/presenterdto"
	"github.com/google/uuid"
)

type ZoomedAreaUpdatedApplicationEvent struct {
	Area      presenterdto.AreaPresenterDto      `json:"area"`
	UnitBlock presenterdto.UnitBlockPresenterDto `json:"unitBlock"`
}

func NewZoomedAreaUpdatedApplicationEventTopic(liveGameId uuid.UUID, playerId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-player-%s-zoomed-area-updated", liveGameId, playerId)
}

func NewZoomedAreaUpdatedApplicationEvent(areaPresenterDto presenterdto.AreaPresenterDto, unitBlockPresenterDto presenterdto.UnitBlockPresenterDto) ZoomedAreaUpdatedApplicationEvent {
	return ZoomedAreaUpdatedApplicationEvent{
		Area:      areaPresenterDto,
		UnitBlock: unitBlockPresenterDto,
	}
}
