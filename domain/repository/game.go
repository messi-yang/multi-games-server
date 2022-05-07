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

type gameModelImpl struct {
}

var gameModel GameRepository

func GetGameRepository() GameRepository {
	if gameModel == nil {
		gameModel = &gameModelImpl{}
		return gameModel
	} else {
		return gameModel
	}
}

func (gmi *gameModelImpl) CreateGameUnitMatrix() {
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

func (gmi *gameModelImpl) GetGameUnitMatrix() *valueobject.GameUnitMatrix {
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

func (gmi *gameModelImpl) GetGameSize() *valueobject.GameSize {
	size := config.GetConfig().GetGameSize()

	return &valueobject.GameSize{Width: size, Height: size}
}
