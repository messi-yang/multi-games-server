package memory

import (
	"math/rand"

	"github.com/DumDumGeniuss/game-of-liberty-computer/config"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/valueobject"
)

type GameRoomMemoryRepository interface {
	GetGameUnitMatrix() *valueobject.GameUnitMatrix
	GetMapSize() *valueobject.MapSize
}

type gameRoomMemoryRepositoryImpl struct {
}

var gameRoomMemoryRepository GameRoomMemoryRepository

func NewGameRoomMemoryRepository() GameRoomMemoryRepository {
	return &gameRoomMemoryRepositoryImpl{}
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
