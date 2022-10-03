package repository

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/aggregate"
	"github.com/google/uuid"
)

type GameRoomRepository interface {
	Add(aggregate.GameRoom) error
	Get(gameId uuid.UUID) (gameRoom aggregate.GameRoom, err error)
	Update(gameId uuid.UUID, gameRoom aggregate.GameRoom) error
	GetAll() []aggregate.GameRoom

	ReadLockAccess(gameId uuid.UUID) (rUnlocker func(), err error)
	LockAccess(gameId uuid.UUID) (unlocker func(), err error)
}
