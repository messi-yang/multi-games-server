package gameunitsrevivedevent

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/coordinatedto"
	"github.com/google/uuid"
)

type Event interface {
	Publish(gameId uuid.UUID, coordinates []coordinatedto.DTO, updatedAt time.Time)
	Subscribe(gameId uuid.UUID, callback func(coordinates []coordinatedto.DTO, updatedAt time.Time)) (unsubscriber func())
}
