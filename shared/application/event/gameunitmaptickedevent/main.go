package gameunitmaptickedevent

import (
	"github.com/google/uuid"
)

type Event interface {
	Publish(gameId uuid.UUID)
	Subscribe(gameId uuid.UUID, callback func()) (unsubscriber func())
}
