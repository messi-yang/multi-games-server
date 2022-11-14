package applicationevent

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/domain/livegame/port/dto"
	"github.com/google/uuid"
)

type ZoomedAreaUpdatedApplicationEvent struct {
	Area      dto.AreaDto      `json:"area"`
	UnitBlock dto.UnitBlockDto `json:"unitBlock"`
}

func NewZoomedAreaUpdatedApplicationEventTopic(liveGameId uuid.UUID, playerId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-player-%s-zoomed-area-updated", liveGameId, playerId)
}

func NewZoomedAreaUpdatedApplicationEvent(areaDto dto.AreaDto, unitBlockDto dto.UnitBlockDto) ZoomedAreaUpdatedApplicationEvent {
	return ZoomedAreaUpdatedApplicationEvent{
		Area:      areaDto,
		UnitBlock: unitBlockDto,
	}
}
