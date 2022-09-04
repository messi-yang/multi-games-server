package gameunitsrevivedevent

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/coordinatedto"
	"github.com/google/uuid"
)

type Event interface {
	Publish(gameId uuid.UUID, coordinates []coordinatedto.Dto)
	Subscribe(gameId uuid.UUID, callback func(coordinates []coordinatedto.Dto)) (unsubscriber func())
}
