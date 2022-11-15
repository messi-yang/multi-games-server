package applicationevent

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/uipresenter/uidto"
	"github.com/google/uuid"
)

type ZoomedAreaUpdatedApplicationEvent struct {
	Area      uidto.AreaUiDto      `json:"area"`
	UnitBlock uidto.UnitBlockUiDto `json:"unitBlock"`
}

func NewZoomedAreaUpdatedApplicationEventTopic(liveGameId uuid.UUID, playerId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-player-%s-zoomed-area-updated", liveGameId, playerId)
}

func NewZoomedAreaUpdatedApplicationEvent(areaUiDto uidto.AreaUiDto, unitBlockUiDto uidto.UnitBlockUiDto) ZoomedAreaUpdatedApplicationEvent {
	return ZoomedAreaUpdatedApplicationEvent{
		Area:      areaUiDto,
		UnitBlock: unitBlockUiDto,
	}
}
