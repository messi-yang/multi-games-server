package gameroomrepository

import (
	"errors"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/aggregate"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/entity"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
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
	GetAll() []aggregate.GameRoom

	AddPlayer(gameId uuid.UUID, player entity.Player) error
	RemovePlayer(gameId uuid.UUID, playerId uuid.UUID) error

	UpdateUnits(uuid.UUID, []valueobject.Coordinate, []valueobject.Unit) (err error)
	UpdateUnitMap(uuid.UUID, *valueobject.UnitMap) (err error)

	GetLastTickedAt(gameId uuid.UUID) (time.Time, error)
	UpdateLastTickedAt(uuid.UUID, time.Time) (err error)

	ReadLockAccess(gameId uuid.UUID) (rUnlocker func(), err error)
	LockAccess(gameId uuid.UUID) (unlocker func(), err error)
}
