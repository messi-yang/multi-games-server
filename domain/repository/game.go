package repository

import (
	"math/rand"

	"github.com/DumDumGeniuss/game-of-liberty-computer/config"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/valueobject"
)

type GameRepository interface {
	GetGameUnitMatrix() *valueobject.GameUnitMatrix
	GetMapSize() *valueobject.MapSize
}

type gameRepositoryImpl struct {
}

var gameRepository GameRepository

func GetGameRepository() GameRepository {
	if gameRepository == nil {
		gameRepository = &gameRepositoryImpl{}
		return gameRepository
	} else {
		return gameRepository
	}
}

func (gmi *gameRepositoryImpl) GetGameUnitMatrix() *valueobject.GameUnitMatrix {
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

func (gmi *gameRepositoryImpl) GetMapSize() *valueobject.MapSize {
	size := config.GetConfig().GetMapSize()

	mapSize := valueobject.NewMapSize(size, size)
	return &mapSize
}
