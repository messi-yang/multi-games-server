package gameroomrepository

import (
	"errors"

	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/aggregate"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/valueobject"
	"github.com/google/uuid"
)

var (
	ErrGameRoomNotFound = errors.New("The game room with given id not found.")
)

type GameRoomRepository interface {
	Add(aggregate.GameRoom) error
	UpdateUnit(uuid.UUID, valueobject.Coordinate, valueobject.Unit) error
	UpdateUnitMatrix(uuid.UUID, [][]valueobject.Unit) error
	Get(uuid.UUID) (aggregate.GameRoom, error)
	GetAll() []aggregate.GameRoom
}
