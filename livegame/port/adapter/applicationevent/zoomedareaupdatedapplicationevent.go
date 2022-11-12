package applicationevent

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/port/dto"
	"github.com/google/uuid"
)

type ZoomedAreaUpdatedApplicationEvent struct {
	Area      dto.AreaDto      `json:"area"`
	UnitBlock dto.UnitBlockDto `json:"unitBlock"`
}

func NewZoomedAreaUpdatedApplicationEventTopic(gameId uuid.UUID, playerId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-player-%s-zoomed-area-updated", gameId, playerId)
}

func NewZoomedAreaUpdatedApplicationEvent(areaDto dto.AreaDto, unitBlockDto dto.UnitBlockDto) ZoomedAreaUpdatedApplicationEvent {
	return ZoomedAreaUpdatedApplicationEvent{
		Area:      areaDto,
		UnitBlock: unitBlockDto,
	}
}
