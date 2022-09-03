package gameunitmaptickedevent

import (
	"time"

	"github.com/google/uuid"
)

type Event interface {
	Publish(gameId uuid.UUID, updatedAt time.Time)
	Subscribe(gameId uuid.UUID, callback func(updatedAt time.Time)) (unsubscriber func())
}
