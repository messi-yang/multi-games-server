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
	config config.Config
}

func CreateGameUnitsModel() {
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

func GetGameUnitsModel() *GameUnitsModel {
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

func GetGameSizeModel() *GameSizeModel {
	size := config.GetConfig().GetGameSize()

	return &GameSizeModel{Width: size, Height: size}
}
