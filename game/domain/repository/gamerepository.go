package repository

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/game/domain/aggregate"
	"github.com/dum-dum-genius/game-of-liberty-computer/game/domain/valueobject"
)

type GameRepository interface {
	Add(aggregate.Game) error
	Get(gameId valueobject.GameId) (game aggregate.Game, err error)
	Update(gameId valueobject.GameId, game aggregate.Game) error
	GetAll() []aggregate.Game

	ReadLockAccess(gameId valueobject.GameId) (rUnlocker func(), err error)
	LockAccess(gameId valueobject.GameId) (unlocker func(), err error)
}
