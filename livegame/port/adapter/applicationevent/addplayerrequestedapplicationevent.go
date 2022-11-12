package applicationevent

import (
	"fmt"

	"github.com/google/uuid"
)

type AddPlayerRequestedApplicationEvent struct {
	PlayerId uuid.UUID `json:"playerId"`
}

func NewAddPlayerRequestedApplicationEventTopic(liveGameId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-add-player-requested", liveGameId)
}

func NewAddPlayerRequestedApplicationEvent(playerId uuid.UUID) AddPlayerRequestedApplicationEvent {
	return AddPlayerRequestedApplicationEvent{
		PlayerId: playerId,
	}
}
