package applicationevent

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/game/port/dto"
	"github.com/google/uuid"
)

type AddPlayerRequestedApplicationEvent struct {
	PlayerId dto.PlayerIdDto `json:"playerId"`
}

func NewAddPlayerRequestedApplicationEventTopic(gameId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-add-player-requested", gameId)
}

func NewAddPlayerRequestedApplicationEvent(playerIdDto dto.PlayerIdDto) AddPlayerRequestedApplicationEvent {
	return AddPlayerRequestedApplicationEvent{
		PlayerId: playerIdDto,
	}
}
