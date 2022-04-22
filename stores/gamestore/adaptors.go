package gamestore

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/daos/gamedao"
)

func convertGameFieldToGameUnits(gameField *gamedao.GameField) *GameUnits {
	width := len(*gameField)
	gameUnits := make(GameUnits, width)
	for i := 0; i < width; i += 1 {
		height := len((*gameField)[i])
		gameUnits[i] = make([]GameUnit, height)
		for j := 0; j < height; j += 1 {

			gameUnits[i][j] = GameUnit{
				Alive: (*gameField)[i][j].Alive,
				Age:   (*gameField)[i][j].Age,
			}
		}
	}

	return &gameUnits
}

func convertGameFieldSizeToGameSize(gameFieldSize *gamedao.GameFieldSize) *GameSize {
	return &GameSize{
		Width:  gameFieldSize.Width,
		Height: gameFieldSize.Height,
	}
}
