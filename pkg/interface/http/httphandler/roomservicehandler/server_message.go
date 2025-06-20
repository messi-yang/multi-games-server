package roomservicehandler

import (
	"fmt"

	"github.com/google/uuid"
)

func newRoomMessageChannel(roomIdDto uuid.UUID) string {
	return fmt.Sprintf("ROOM_%s_CHANNEL", roomIdDto)
}

type roomMessage[T any] struct {
	SenderId    uuid.UUID `json:"senderId"`
	ServerEvent T         `json:"serverEvent"`
}

func newPlayerMessageChannel(roomIdDto uuid.UUID, playerIdDto uuid.UUID) string {
	return fmt.Sprintf("ROOM_%s_PLAYER_%s_CHANNEL", roomIdDto, playerIdDto)
}

type playerMessage[T any] struct {
	ServerEvent T `json:"serverEvent"`
}
