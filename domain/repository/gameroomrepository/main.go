package gameroomrepository

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/aggregate"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/valueobject"
	"github.com/google/uuid"
)

type GameRoomRepository interface {
	Create(aggregate.GameRoom) error
	UpdateGameUnit(uuid.UUID, valueobject.Coordinate, valueobject.GameUnit) error
	UpdateGameUnitMatrix(uuid.UUID, [][]valueobject.GameUnit) error
	Get(uuid.UUID) (aggregate.GameRoom, error)
}
