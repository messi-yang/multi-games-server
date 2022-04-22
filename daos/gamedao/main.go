package gamedao

import (
	"math/rand"

	"github.com/DumDumGeniuss/game-of-liberty-computer/config"
	"github.com/DumDumGeniuss/game-of-liberty-computer/models/gamemodel"
)

type GameDao struct {
}

func GetGameDao() *GameDao {
	return &GameDao{}
}

func (gd *GameDao) GetGameField() (*gamemodel.GameField, error) {
	size := config.Config.GAME_SIZE
	gameField := make(gamemodel.GameField, size)

	for i := 0; i < size; i += 1 {
		gameField[i] = make([]gamemodel.GameFieldUnit, size)
		for j := 0; j < size; j += 1 {
			gameField[i][j] = gamemodel.GameFieldUnit{
				Alive: rand.Intn(2) == 0,
				Age:   0,
			}
		}
	}

	return &gameField, nil
}

func (gd *GameDao) GetGameFieldSize() (*gamemodel.GameFieldSize, error) {
	size := config.Config.GAME_SIZE
	return &gamemodel.GameFieldSize{Width: size, Height: size}, nil
}
