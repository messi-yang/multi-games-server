package applicationevent

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/presenter/presenterdto"
	"github.com/google/uuid"
)

type AreaZoomedApplicationEvent struct {
	Area      presenterdto.AreaPresenterDto      `json:"area"`
	UnitBlock presenterdto.UnitBlockPresenterDto `json:"unitBlock"`
}

func NewAreaZoomedApplicationEventTopic(liveGameId uuid.UUID, playerId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-player-%s-area-zoomed", liveGameId, playerId)
}

func NewAreaZoomedApplicationEvent(areaPresenterDto presenterdto.AreaPresenterDto, unitBlockPresenterDto presenterdto.UnitBlockPresenterDto) AreaZoomedApplicationEvent {
	return AreaZoomedApplicationEvent{
		Area:      areaPresenterDto,
		UnitBlock: unitBlockPresenterDto,
	}
}
