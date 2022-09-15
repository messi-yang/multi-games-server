package gameunitsrevivedevent

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"
	"github.com/google/uuid"
)

type Event interface {
	Publish(gameId uuid.UUID, coordinates []valueobject.Coordinate)
	Subscribe(gameId uuid.UUID, callback func(coordinates []valueobject.Coordinate)) (unsubscriber func())
}
