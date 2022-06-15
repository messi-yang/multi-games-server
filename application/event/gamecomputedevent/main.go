package gamecomputedevent

import "github.com/google/uuid"

type GameComputedEvent interface {
	Publish(gameId uuid.UUID)
	Subscribe(gameId uuid.UUID, callback func())
	Unsubscribe(gameId uuid.UUID, callback func())
}
