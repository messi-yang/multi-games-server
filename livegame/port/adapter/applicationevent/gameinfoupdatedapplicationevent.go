package applicationevent

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/port/dto"
	"github.com/google/uuid"
)

type GameInfoUpdatedApplicationEvent struct {
	Dimension dto.DimensionDto `json:"dimension"`
}

func NewGameInfoUpdatedApplicationEventTopic(gameId uuid.UUID, playerId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-player-%s-game-info-updated", gameId, playerId)
}

func NewGameInfoUpdatedApplicationEvent(dimensionDto dto.DimensionDto) GameInfoUpdatedApplicationEvent {
	return GameInfoUpdatedApplicationEvent{
		Dimension: dimensionDto,
	}
}
