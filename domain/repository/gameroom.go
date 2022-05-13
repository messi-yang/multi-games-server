package repository

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/aggregate"
	"github.com/google/uuid"
)

type GameRoomRepository interface {
	GetAll() ([]aggregate.GameRoom, error)
	GetById(uuid.UUID) (aggregate.GameRoom, error)
	Add(aggregate.GameRoom) error
}
