package repository

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/game/domain/aggregate"
	gameValueObject "github.com/dum-dum-genius/game-of-liberty-computer/game/domain/valueobject"
)

type GameRepository interface {
	Add(aggregate.Game) error
	Get(gameValueObject.GameId) (aggregate.Game, error)
	Update(gameValueObject.GameId, aggregate.Game) error
	GetAll() []aggregate.Game

	ReadLockAccess(gameValueObject.GameId) (rUnlocker func(), err error)
	LockAccess(gameValueObject.GameId) (unlocker func(), err error)
}
