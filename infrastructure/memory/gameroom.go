package memory

import (
	"fmt"

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
		gameRoomMemoryRepository = &gameRoomMemoryRepositoryImpl{
			gameRoomMap: make(map[uuid.UUID]aggregate.GameRoom),
		}
		return gameRoomMemoryRepository
	} else {
		return gameRoomMemoryRepository
	}
}

func (gmi *gameRoomMemoryRepositoryImpl) Get(id uuid.UUID) (aggregate.GameRoom, error) {
	fmt.Println(gmi.gameRoomMap)
	return gmi.gameRoomMap[id], nil
}

func (gmi *gameRoomMemoryRepositoryImpl) Add(gameRoom aggregate.GameRoom) error {
	gmi.gameRoomMap[gameRoom.GetGameId()] = gameRoom

	return nil
}
