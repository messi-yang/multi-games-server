package applicationevent

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/game/port/dto"
	"github.com/google/uuid"
)

type GameInfoUpdatedApplicationEvent struct {
	Dimension dto.DimensionDto `json:"dimension"`
}

func NewGameInfoUpdatedApplicationEventTopic(gameId uuid.UUID, playerIdDto dto.PlayerIdDto) string {
	return fmt.Sprintf("game-room-%s-player-%s-game-info-updated", gameId, playerIdDto.Value)
}

func NewGameInfoUpdatedApplicationEvent(dimensionDto dto.DimensionDto) GameInfoUpdatedApplicationEvent {
	return GameInfoUpdatedApplicationEvent{
		Dimension: dimensionDto,
	}
}
