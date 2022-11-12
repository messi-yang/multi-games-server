package repository

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/domain/aggregate"
	liveGameValueObject "github.com/dum-dum-genius/game-of-liberty-computer/livegame/domain/valueobject"
)

type LiveGameRepository interface {
	Add(aggregate.LiveGame) error
	Get(liveGameValueObject.LiveGameId) (aggregate.LiveGame, error)
	Update(liveGameValueObject.LiveGameId, aggregate.LiveGame) error
	GetAll() []aggregate.LiveGame

	ReadLockAccess(liveGameValueObject.LiveGameId) (rUnlocker func(), err error)
	LockAccess(liveGameValueObject.LiveGameId) (unlocker func(), err error)
}
