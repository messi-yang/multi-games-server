package applicationevent

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/game/port/dto"
	"github.com/google/uuid"
)

type AreaZoomedApplicationEvent struct {
	Area      dto.AreaDto      `json:"area"`
	UnitBlock dto.UnitBlockDto `json:"unitBlock"`
}

func NewAreaZoomedApplicationEventTopic(gameId uuid.UUID, playerId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-player-%s-area-zoomed", gameId, playerId)
}

func NewAreaZoomedApplicationEvent(areaDto dto.AreaDto, unitBlockDto dto.UnitBlockDto) AreaZoomedApplicationEvent {
	return AreaZoomedApplicationEvent{
		Area:      areaDto,
		UnitBlock: unitBlockDto,
	}
}
