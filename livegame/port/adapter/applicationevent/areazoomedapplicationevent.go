package applicationevent

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/port/dto"
	"github.com/google/uuid"
)

type AreaZoomedApplicationEvent struct {
	Area      dto.AreaDto      `json:"area"`
	UnitBlock dto.UnitBlockDto `json:"unitBlock"`
}

func NewAreaZoomedApplicationEventTopic(liveGameId uuid.UUID, playerId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-player-%s-area-zoomed", liveGameId, playerId)
}

func NewAreaZoomedApplicationEvent(areaDto dto.AreaDto, unitBlockDto dto.UnitBlockDto) AreaZoomedApplicationEvent {
	return AreaZoomedApplicationEvent{
		Area:      areaDto,
		UnitBlock: unitBlockDto,
	}
}
