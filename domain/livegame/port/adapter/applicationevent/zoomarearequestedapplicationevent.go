package applicationevent

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/presenter/presenterdto"
	"github.com/google/uuid"
)

type ZoomAreaRequestedApplicationEvent struct {
	PlayerId uuid.UUID                     `json:"playerId"`
	Area     presenterdto.AreaPresenterDto `json:"area"`
}

func NewZoomAreaRequestedApplicationEventTopic(liveGameId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-zoom-area-requested", liveGameId)
}

func NewZoomAreaRequestedApplicationEvent(playerId uuid.UUID, areaPresenterDto presenterdto.AreaPresenterDto) ZoomAreaRequestedApplicationEvent {
	return ZoomAreaRequestedApplicationEvent{
		PlayerId: playerId,
		Area:     areaPresenterDto,
	}
}
