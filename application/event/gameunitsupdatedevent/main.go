package gameunitsupdatedevent

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/coordinatedto"
	"github.com/google/uuid"
)

type UnitsUpdatedEvent interface {
	Publish(gameId uuid.UUID, coordinates []coordinatedto.CoordinateDTO, updatedAt time.Time)
	Subscribe(gameId uuid.UUID, callback func(coordinates []coordinatedto.CoordinateDTO, updatedAt time.Time)) (unsubscriber func())
}
