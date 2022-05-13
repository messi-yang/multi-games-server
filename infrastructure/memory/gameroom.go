package memory

import (
	"math/rand"

	"github.com/DumDumGeniuss/game-of-liberty-computer/config"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/aggregate"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/valueobject"
	"github.com/google/uuid"
)

type GameRoomMemoryRepository interface {
	GetGameUnitMatrix() *valueobject.GameUnitMatrix
	GetMapSize() *valueobject.MapSize
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

func (gmi *gameRoomMemoryRepositoryImpl) GetGameUnitMatrix() *valueobject.GameUnitMatrix {
	size := config.GetConfig().GetMapSize()

	gameUnitsEntity := make(valueobject.GameUnitMatrix, size)
	for i := 0; i < size; i += 1 {
		gameUnitsEntity[i] = make([]valueobject.GameUnit, size)
		for j := 0; j < size; j += 1 {
			gameUnitsEntity[i][j] = valueobject.GameUnit{
				Alive: rand.Intn(2) == 0,
				Age:   0,
			}
		}
	}

	return &gameUnitsEntity
}

func (gmi *gameRoomMemoryRepositoryImpl) GetMapSize() *valueobject.MapSize {
	size := config.GetConfig().GetMapSize()

	mapSize := valueobject.NewMapSize(size, size)
	return &mapSize
}
