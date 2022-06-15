package gameunitsupdatedevent

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/valueobject"
	"github.com/google/uuid"
)

type GameUnitsUpdatedEvent interface {
	Publish(gameId uuid.UUID, coordinates []valueobject.Coordinate)
	Subscribe(gameId uuid.UUID, callback func(coordinates []valueobject.Coordinate))
	Unsubscribe(gameId uuid.UUID, callback func(coordinates []valueobject.Coordinate))
}
