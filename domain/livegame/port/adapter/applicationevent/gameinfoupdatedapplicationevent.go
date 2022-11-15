package applicationevent

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/presenter/presenterdto"
	"github.com/google/uuid"
)

type GameInfoUpdatedApplicationEvent struct {
	Dimension presenterdto.DimensionPresenterDto `json:"dimension"`
}

func NewGameInfoUpdatedApplicationEventTopic(liveGameId uuid.UUID, playerId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-player-%s-game-info-updated", liveGameId, playerId)
}

func NewGameInfoUpdatedApplicationEvent(dimensionPresenterDto presenterdto.DimensionPresenterDto) GameInfoUpdatedApplicationEvent {
	return GameInfoUpdatedApplicationEvent{
		Dimension: dimensionPresenterDto,
	}
}
