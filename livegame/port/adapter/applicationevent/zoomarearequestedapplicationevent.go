package applicationevent

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/port/dto"
	"github.com/google/uuid"
)

type ZoomAreaRequestedApplicationEvent struct {
	PlayerId uuid.UUID   `json:"playerId"`
	Area     dto.AreaDto `json:"area"`
}

func NewZoomAreaRequestedApplicationEventTopic(liveGameId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-zoom-area-requested", liveGameId)
}

func NewZoomAreaRequestedApplicationEvent(playerId uuid.UUID, areaDto dto.AreaDto) ZoomAreaRequestedApplicationEvent {
	return ZoomAreaRequestedApplicationEvent{
		PlayerId: playerId,
		Area:     areaDto,
	}
}
