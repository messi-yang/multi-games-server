package applicationevent

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/game/port/dto"
	"github.com/google/uuid"
)

type RemovePlayerRequestedApplicationEvent struct {
	PlayerId dto.PlayerIdDto `json:"playerId"`
}

func NewRemovePlayerRequestedApplicationEventTopic(gameId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-remove-player-requested", gameId)
}

func NewRemovePlayerRequestedApplicationEvent(playerIdDto dto.PlayerIdDto) RemovePlayerRequestedApplicationEvent {
	return RemovePlayerRequestedApplicationEvent{
		PlayerId: playerIdDto,
	}
}
