package gamedao

import (
	"math/rand"

	"github.com/DumDumGeniuss/game-of-liberty-computer/config"
	"github.com/DumDumGeniuss/game-of-liberty-computer/models/gamemodel"
)

type dao struct {
}

func GetDao() *dao {
	return &dao{}
}

func (d *dao) GetGameField() (*GameField, error) {
	size := config.Config.GAME_SIZE

	// TODO - Get gameFieldEntity from real storage
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

	gameField := make(GameField, size)
	for i := 0; i < size; i += 1 {
		gameField[i] = make([]GameFieldUnit, size)
		for j := 0; j < size; j += 1 {

			gameField[i][j] = GameFieldUnit{
				Alive: gameFieldEntity[i][j].Alive,
				Age:   gameFieldEntity[i][j].Age,
			}
		}
	}

	return &gameField, nil
}

func (d *dao) GetGameFieldSize() (*GameFieldSize, error) {
	size := config.Config.GAME_SIZE
	gameFieldSizeEntity := gamemodel.GameFieldSizeEntity{Width: size, Height: size}

	gameFieldSize := GameFieldSize{
		Width:  gameFieldSizeEntity.Width,
		Height: gameFieldSizeEntity.Height,
	}
	return &gameFieldSize, nil
}
