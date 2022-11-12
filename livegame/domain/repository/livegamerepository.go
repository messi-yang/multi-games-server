package repository

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/domain/aggregate"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/domain/valueobject"
)

type LiveGameRepository interface {
	Add(aggregate.LiveGame) error
	Get(valueobject.GameId) (aggregate.LiveGame, error)
	Update(valueobject.GameId, aggregate.LiveGame) error
	GetAll() []aggregate.LiveGame

	ReadLockAccess(valueobject.GameId) (rUnlocker func(), err error)
	LockAccess(valueobject.GameId) (unlocker func(), err error)
}
