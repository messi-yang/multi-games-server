package gameroomrepository

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/aggregate"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"
	"github.com/google/uuid"
)

var (
	ErrGameRoomNotFound = errors.New("the game room with given id not found")
)

type GameRoomRepository interface {
	Add(aggregate.GameRoom) error
	UpdateUnit(uuid.UUID, valueobject.Coordinate, valueobject.Unit) error
	UpdateUnitMap(uuid.UUID, valueobject.UnitMap) error
	Get(uuid.UUID) (aggregate.GameRoom, error)
	GetAll() []aggregate.GameRoom
}
