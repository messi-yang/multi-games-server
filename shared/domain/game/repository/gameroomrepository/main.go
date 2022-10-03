package gameroomrepository

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/aggregate"
	"github.com/google/uuid"
)

var (
	ErrGameRoomNotFound       = errors.New("the game room with the id was not found")
	ErrGameRoomLockerNotFound = errors.New("the game room locker for with the id was not found")
	ErrPlayerAlreadyExists    = errors.New("the player with the given id alredy exists in the game room")
)

type Repository interface {
	Add(aggregate.GameRoom) error
	Get(gameId uuid.UUID) (gameRoom aggregate.GameRoom, err error)
	Update(gameId uuid.UUID, gameRoom aggregate.GameRoom) error
	GetAll() []aggregate.GameRoom

	ReadLockAccess(gameId uuid.UUID) (rUnlocker func(), err error)
	LockAccess(gameId uuid.UUID) (unlocker func(), err error)
}
