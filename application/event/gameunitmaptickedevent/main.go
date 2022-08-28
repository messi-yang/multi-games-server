package gameunitmaptickedevent

import (
	"time"

	"github.com/google/uuid"
)

type GameUnitMapTickedEvent interface {
	Publish(gameId uuid.UUID, updatedAt time.Time)
	Subscribe(gameId uuid.UUID, callback func(updatedAt time.Time)) (unsubscriber func())
}
