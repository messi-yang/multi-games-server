package worldjourneyhandler

import (
	"fmt"

	"github.com/google/uuid"
)

func newWorldMessageChannel(worldIdDto uuid.UUID) string {
	return fmt.Sprintf("WORLD_%s_CHANNEL", worldIdDto)
}

type worldMessage[T any] struct {
	SenderId    uuid.UUID `json:"senderId"`
	ServerEvent T         `json:"serverEvent"`
}

func newPlayerMessageChannel(worldIdDto uuid.UUID, playerIdDto uuid.UUID) string {
	return fmt.Sprintf("WORLD_%s_PLAYER_%s_CHANNEL", worldIdDto, playerIdDto)
}

type playerMessage[T any] struct {
	ServerEvent T `json:"serverEvent"`
}
