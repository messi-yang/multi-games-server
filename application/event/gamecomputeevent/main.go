package gamecomputeevent

import "github.com/google/uuid"

type GameComputeEvent interface {
	Publish(gameId uuid.UUID)
	Subscribe(gameId uuid.UUID, callback func())
	Unsubscribe(gameId uuid.UUID, callback func())
}
