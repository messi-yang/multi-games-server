package gamemodel

import (
	"math/rand"

	"github.com/DumDumGeniuss/game-of-liberty-computer/config"
)

type GameModel interface {
	CreateGameUnitsModel()
	GetGameUnitsModel() *GameUnitsModel
	GetGameSizeModel() *GameSizeModel
}

type gameModelImpl struct {
}

var gameModel GameModel

func GetGameModel() GameModel {
	if gameModel == nil {
		gameModel = &gameModelImpl{}
		return gameModel
	} else {
		return gameModel
	}
}

func (gmi *gameModelImpl) CreateGameUnitsModel() {
	size := config.GetConfig().GetGameSize()

	gameUnitsEntity := make(GameUnitsModel, size)
	for i := 0; i < size; i += 1 {
		gameUnitsEntity[i] = make([]GameUnitModel, size)
		for j := 0; j < size; j += 1 {
			gameUnitsEntity[i][j] = GameUnitModel{
				Alive: rand.Intn(2) == 0,
				Age:   0,
			}
		}
	}

	// TODO - Insert into whatever storage we use.
}

func (gmi *gameModelImpl) GetGameUnitsModel() *GameUnitsModel {
	size := config.GetConfig().GetGameSize()

	gameUnitsEntity := make(GameUnitsModel, size)
	for i := 0; i < size; i += 1 {
		gameUnitsEntity[i] = make([]GameUnitModel, size)
		for j := 0; j < size; j += 1 {
			gameUnitsEntity[i][j] = GameUnitModel{
				Alive: rand.Intn(2) == 0,
				Age:   0,
			}
		}
	}

	return &gameUnitsEntity
}

func (gmi *gameModelImpl) GetGameSizeModel() *GameSizeModel {
	size := config.GetConfig().GetGameSize()

	return &GameSizeModel{Width: size, Height: size}
}
