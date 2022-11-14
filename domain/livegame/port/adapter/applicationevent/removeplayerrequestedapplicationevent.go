package applicationevent

import (
	"fmt"

	"github.com/google/uuid"
)

type RemovePlayerRequestedApplicationEvent struct {
	PlayerId uuid.UUID `json:"playerId"`
}

func NewRemovePlayerRequestedApplicationEventTopic(liveGameId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-remove-player-requested", liveGameId)
}

func NewRemovePlayerRequestedApplicationEvent(playerId uuid.UUID) RemovePlayerRequestedApplicationEvent {
	return RemovePlayerRequestedApplicationEvent{
		PlayerId: playerId,
	}
}
