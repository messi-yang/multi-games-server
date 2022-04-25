package gamestore

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/daos/gamedao"
	"github.com/DumDumGeniuss/ggol"
)

type GameStore interface {
	GetGameUnitsInArea(area *GameArea) (*[][]*GameUnit, error)
	GetGameSize() *GameSize
}

type gameStoreImplement struct {
	gameOfLiberty ggol.Game[GameUnit]
	gameDAO       gamedao.GameDAO
}

var gameStore GameStore = nil

func GetGameStore(gameDAO gamedao.GameDAO) GameStore {
	if gameStore == nil {
		initialUnit := GameUnit{
			Alive: false,
			Age:   0,
		}
		gameFieldSize, _ := gameDAO.GetGameSize()
		ggolSize := convertGameSizeToGgolSize(gameFieldSize)
		newGameOfLiberty, _ := ggol.NewGame(
			ggolSize,
			&initialUnit,
		)

		newGameOfLiberty.SetNextUnitGenerator(gameNextUnitGenerator)

		gameField, _ := gameDAO.GetGameField()
		gameUnits := convertGameFieldToGameUnits(gameField)

		for x := 0; x < ggolSize.Width; x += 1 {
			for y := 0; y < ggolSize.Height; y += 1 {
				gameFieldUnit := &(*gameUnits)[x][y]
				coord := &ggol.Coordinate{X: x, Y: y}
				newGameOfLiberty.SetUnit(coord, gameFieldUnit)
			}
		}

		newGameStore := &gameStoreImplement{
			gameOfLiberty: newGameOfLiberty,
			gameDAO:       gameDAO,
		}
		gameStore = newGameStore
		return gameStore
	} else {
		return gameStore
	}
}

func (gsi *gameStoreImplement) GetGameUnitsInArea(area *GameArea) (*[][]*GameUnit, error) {
	ggolArea := convertGameAreaToGgolArea(area)
	units, err := gsi.gameOfLiberty.GetUnitsInArea(ggolArea)
	if err != nil {
		return nil, err
	}
	return &units, nil
}

func (gsi *gameStoreImplement) GetGameSize() *GameSize {
	return convertGgolSizeToGameSize(gsi.gameOfLiberty.GetSize())
}
