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
	Create(aggregate.GameRoom) error
	UpdateGameUnit(uuid.UUID, valueobject.Coordinate, valueobject.GameUnit) error
	UpdateGameUnitMatrix(uuid.UUID, [][]valueobject.GameUnit) error
	Get(uuid.UUID) (aggregate.GameRoom, error)
}
