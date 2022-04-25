package gameprovider

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/daos/gamedao"
	"github.com/DumDumGeniuss/ggol"
)

func convertGameUnitsFromGameDAOToGameUnits(gameUnitsFromGameDAO *gamedao.GameUnits) *GameUnits {
	width := len(*gameUnitsFromGameDAO)
	gameUnits := make(GameUnits, width)
	for i := 0; i < width; i += 1 {
		height := len((*gameUnitsFromGameDAO)[i])
		gameUnits[i] = make([]GameUnit, height)
		for j := 0; j < height; j += 1 {

			gameUnits[i][j] = GameUnit{
				Alive: (*gameUnitsFromGameDAO)[i][j].Alive,
				Age:   (*gameUnitsFromGameDAO)[i][j].Age,
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

func convertGameSizeToGgolSize(gameSize *gamedao.GameSize) *ggol.Size {
	return &ggol.Size{
		Width:  gameSize.Width,
		Height: gameSize.Height,
	}
}

func convertGameAreaToGgolArea(gameArea *GameArea) *ggol.Area {
	return &ggol.Area{
		From: ggol.Coordinate(gameArea.From),
		To:   ggol.Coordinate(gameArea.To),
	}
}
