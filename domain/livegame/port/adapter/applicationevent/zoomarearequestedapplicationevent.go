package applicationevent

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/uipresenter/uidto"
	"github.com/google/uuid"
)

type ZoomAreaRequestedApplicationEvent struct {
	PlayerId uuid.UUID       `json:"playerId"`
	Area     uidto.AreaUiDto `json:"area"`
}

func NewZoomAreaRequestedApplicationEventTopic(liveGameId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-zoom-area-requested", liveGameId)
}

func NewZoomAreaRequestedApplicationEvent(playerId uuid.UUID, areaUiDto uidto.AreaUiDto) ZoomAreaRequestedApplicationEvent {
	return ZoomAreaRequestedApplicationEvent{
		PlayerId: playerId,
		Area:     areaUiDto,
	}
}
