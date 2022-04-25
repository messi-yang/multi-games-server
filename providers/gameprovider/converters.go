package gameprovider

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/daos/gamedao"
	"github.com/DumDumGeniuss/ggol"
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

func convertGgolSizeToGameSize(ggolSize *ggol.Size) *GameSize {
	return &GameSize{
		Width:  ggolSize.Width,
		Height: ggolSize.Height,
	}
}

func convertGameSizeToGgolSize(gameFieldSize *gamedao.GameSize) *ggol.Size {
	return &ggol.Size{
		Width:  gameFieldSize.Width,
		Height: gameFieldSize.Height,
	}
}

func convertGameAreaToGgolArea(gameArea *GameArea) *ggol.Area {
	return &ggol.Area{
		From: ggol.Coordinate(gameArea.From),
		To:   ggol.Coordinate(gameArea.To),
	}
}
