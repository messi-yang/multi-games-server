package gamedao

import (
	"math/rand"

	"github.com/DumDumGeniuss/game-of-liberty-computer/config"
	"github.com/DumDumGeniuss/game-of-liberty-computer/models/gamemodel"
)

func createGameFieldEntityInStorage() {
	size := config.Config.GAME_SIZE

	gameFieldEntity := make(gamemodel.GameFieldEntity, size)
	for i := 0; i < size; i += 1 {
		gameFieldEntity[i] = make([]gamemodel.GameFieldUnitEntity, size)
		for j := 0; j < size; j += 1 {
			gameFieldEntity[i][j] = gamemodel.GameFieldUnitEntity{
				Alive: rand.Intn(2) == 0,
				Age:   0,
			}
		}
	}

	// TODO - Insert into whatever storage we use.
}

func getGameFieldEntityFromStorage() *gamemodel.GameFieldEntity {
	size := config.Config.GAME_SIZE

	gameFieldEntity := make(gamemodel.GameFieldEntity, size)
	for i := 0; i < size; i += 1 {
		gameFieldEntity[i] = make([]gamemodel.GameFieldUnitEntity, size)
		for j := 0; j < size; j += 1 {
			gameFieldEntity[i][j] = gamemodel.GameFieldUnitEntity{
				Alive: rand.Intn(2) == 0,
				Age:   0,
			}
		}
	}

	return &gameFieldEntity
}

func getGameFieldSizeEntityFromStorage() *gamemodel.GameFieldSizeEntity {
	size := config.Config.GAME_SIZE

	return &gamemodel.GameFieldSizeEntity{Width: size, Height: size}
}
