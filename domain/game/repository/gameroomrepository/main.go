package gameroomrepository

import (
	"errors"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/aggregate"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"
	"github.com/google/uuid"
)

var (
	ErrGameRoomNotFound       = errors.New("the game room with the id was not found")
	ErrGameRoomLockerNotFound = errors.New("the game room locker for with the id was not found")
)

type GameRoomRepository interface {
	Add(aggregate.GameRoom) error
	UpdateUnits(uuid.UUID, []valueobject.Coordinate, []valueobject.Unit) (err error)
	UpdateUnitMap(uuid.UUID, valueobject.UnitMap) (err error)
	UpdateLastTickedAt(uuid.UUID, time.Time) (err error)
	Get(uuid.UUID) (gameRoom aggregate.GameRoom, receivedAt time.Time, err error)
	GetAll() []aggregate.GameRoom
	GetLastTickedAt(uuid.UUID) (time.Time, error)
	ReadLockAccess(uuid.UUID) (rUnlocker func(), err error)
	LockAccess(uuid.UUID) (unlocker func(), err error)
}
