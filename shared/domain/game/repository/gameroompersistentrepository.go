package repository

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/aggregate"
	"github.com/google/uuid"
)

type GameRoomPersistentRepository interface {
	Add(aggregate.GameRoom) error
	Get(gameId uuid.UUID) (gameRoom aggregate.GameRoom, err error)
	Update(gameId uuid.UUID, gameRoom aggregate.GameRoom) error
}
