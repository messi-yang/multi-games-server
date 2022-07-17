package coordinatesupdatedevent

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/dto/coordinatedto"
	"github.com/google/uuid"
)

type CoordinatesUpdatedEvent interface {
	Publish(gameId uuid.UUID, coordinates []coordinatedto.CoordinateDTO)
	Subscribe(gameId uuid.UUID, callback func(coordinates []coordinatedto.CoordinateDTO)) (unsubscriber func())
}
