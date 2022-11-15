package applicationevent

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/uipresenter/uidto"
	"github.com/google/uuid"
)

type AreaZoomedApplicationEvent struct {
	Area      uidto.AreaUiDto      `json:"area"`
	UnitBlock uidto.UnitBlockUiDto `json:"unitBlock"`
}

func NewAreaZoomedApplicationEventTopic(liveGameId uuid.UUID, playerId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-player-%s-area-zoomed", liveGameId, playerId)
}

func NewAreaZoomedApplicationEvent(areaUiDto uidto.AreaUiDto, unitBlockUiDto uidto.UnitBlockUiDto) AreaZoomedApplicationEvent {
	return AreaZoomedApplicationEvent{
		Area:      areaUiDto,
		UnitBlock: unitBlockUiDto,
	}
}
