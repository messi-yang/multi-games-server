package memory

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/aggregate"
	"github.com/google/uuid"
)

type GameRoomMemoryRepository interface {
	Get(uuid.UUID) (aggregate.GameRoom, error)
	Add(aggregate.GameRoom) error
}

type gameRoomMemoryRepositoryImpl struct {
	gameRoomMap map[uuid.UUID]aggregate.GameRoom
}

var gameRoomMemoryRepository GameRoomMemoryRepository

func NewGameRoomMemoryRepository() GameRoomMemoryRepository {
	if gameRoomMemoryRepository == nil {
		return &gameRoomMemoryRepositoryImpl{
			gameRoomMap: make(map[uuid.UUID]aggregate.GameRoom),
		}
	} else {
		return gameRoomMemoryRepository
	}
}

func (gmi *gameRoomMemoryRepositoryImpl) Get(id uuid.UUID) (aggregate.GameRoom, error) {
	return gmi.gameRoomMap[id], nil
}

func (gmi *gameRoomMemoryRepositoryImpl) Add(gameRoom aggregate.GameRoom) error {
	gmi.gameRoomMap[gameRoom.Game.Id] = gameRoom

	return nil
}
