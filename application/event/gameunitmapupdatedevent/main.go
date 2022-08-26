package gameunitmapupdatedevent

import "github.com/google/uuid"

type GameUnitMapUpdatedEvent interface {
	Publish(gameId uuid.UUID)
	Subscribe(gameId uuid.UUID, callback func()) (unsubscriber func())
}
