package applicationevent

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/uipresenter/uidto"
	"github.com/google/uuid"
)

type GameInfoUpdatedApplicationEvent struct {
	Dimension uidto.DimensionUiDto `json:"dimension"`
}

func NewGameInfoUpdatedApplicationEventTopic(liveGameId uuid.UUID, playerId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-player-%s-game-info-updated", liveGameId, playerId)
}

func NewGameInfoUpdatedApplicationEvent(dimensionUiDto uidto.DimensionUiDto) GameInfoUpdatedApplicationEvent {
	return GameInfoUpdatedApplicationEvent{
		Dimension: dimensionUiDto,
	}
}
