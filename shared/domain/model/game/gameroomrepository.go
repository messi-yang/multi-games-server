package game

import (
	"github.com/google/uuid"
)

type GameRoomRepository interface {
	Add(GameRoom) error
	Get(gameId uuid.UUID) (gameRoom GameRoom, err error)
	Update(gameId uuid.UUID, gameRoom GameRoom) error
	GetAll() []GameRoom

	ReadLockAccess(gameId uuid.UUID) (rUnlocker func(), err error)
	LockAccess(gameId uuid.UUID) (unlocker func(), err error)
}
