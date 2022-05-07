package repository

import (
	"math/rand"

	"github.com/DumDumGeniuss/game-of-liberty-computer/config"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/valueobject"
)

type GameRepository interface {
	CreateGameUnitMatrix()
	GetGameUnitMatrix() *valueobject.GameUnitMatrix
	GetGameSize() *valueobject.GameSize
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

func (gmi *gameRepositoryImpl) CreateGameUnitMatrix() {
	size := config.GetConfig().GetGameSize()

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

	// TODO - Insert into whatever storage we use.
}

func (gmi *gameRepositoryImpl) GetGameUnitMatrix() *valueobject.GameUnitMatrix {
	size := config.GetConfig().GetGameSize()

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

func (gmi *gameRepositoryImpl) GetGameSize() *valueobject.GameSize {
	size := config.GetConfig().GetGameSize()

	return &valueobject.GameSize{Width: size, Height: size}
}
