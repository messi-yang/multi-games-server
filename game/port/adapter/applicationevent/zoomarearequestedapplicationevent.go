package applicationevent

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/game/port/dto"
	"github.com/google/uuid"
)

type ZoomAreaRequestedApplicationEvent struct {
	PlayerId dto.PlayerIdDto `json:"playerId"`
	Area     dto.AreaDto     `json:"area"`
}

func NewZoomAreaRequestedApplicationEventTopic(gameId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-zoom-area-requested", gameId)
}

func NewZoomAreaRequestedApplicationEvent(playerIdDto dto.PlayerIdDto, areaDto dto.AreaDto) ZoomAreaRequestedApplicationEvent {
	return ZoomAreaRequestedApplicationEvent{
		PlayerId: playerIdDto,
		Area:     areaDto,
	}
}
