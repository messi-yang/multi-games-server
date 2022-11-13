package repository

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/game/domain/aggregate"
	"github.com/dum-dum-genius/game-of-liberty-computer/game/domain/valueobject"
	gameValueObject "github.com/dum-dum-genius/game-of-liberty-computer/game/domain/valueobject"
)

type GameRepository interface {
	Add(aggregate.Game) (valueobject.GameId, error)
	Get(gameValueObject.GameId) (aggregate.Game, error)
	Update(gameValueObject.GameId, aggregate.Game) error
	GetAll() ([]aggregate.Game, error)

	ReadLockAccess(gameValueObject.GameId) (rUnlocker func(), err error)
	LockAccess(gameValueObject.GameId) (unlocker func(), err error)
}
