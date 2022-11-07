package repository

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/game/domain/aggregate"
	"github.com/google/uuid"
)

type GameRepository interface {
	Add(aggregate.Game) error
	Get(gameId uuid.UUID) (game aggregate.Game, err error)
	Update(gameId uuid.UUID, game aggregate.Game) error
	GetAll() []aggregate.Game

	ReadLockAccess(gameId uuid.UUID) (rUnlocker func(), err error)
	LockAccess(gameId uuid.UUID) (unlocker func(), err error)
}
