package applicationevent

import (
	"fmt"

	"github.com/google/uuid"
)

type AddPlayerRequestedApplicationEvent struct {
	PlayerId uuid.UUID `json:"playerId"`
}

func NewAddPlayerRequestedApplicationEventTopic(gameId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-add-player-requested", gameId)
}

func NewAddPlayerRequestedApplicationEvent(playerId uuid.UUID) AddPlayerRequestedApplicationEvent {
	return AddPlayerRequestedApplicationEvent{
		PlayerId: playerId,
	}
}
