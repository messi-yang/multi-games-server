package gamestore

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/daos/gamedao"
	"github.com/DumDumGeniuss/ggol"
)

type GameStore interface {
	GetGameUnitsInArea(area *GameArea) (*[][]*GameUnit, error)
	GetGameFieldSize() *GameSize
	InitializeGame()
}

type gameStoreImplement struct {
	gameOfLiberty ggol.Game[GameUnit]
	gameDao       gamedao.GameDAO
}

var Store GameStore = &gameStoreImplement{
	gameOfLiberty: nil,
	gameDao:       gamedao.DAO,
}

func (gsi *gameStoreImplement) InitializeGame() {
	initialUnit := GameUnit{
		Alive: false,
		Age:   0,
	}
	gameFieldSize, _ := gsi.gameDao.GetGameFieldSize()
	ggolSize := convertGameGameFieldSizeToGgolSize(gameFieldSize)
	gsi.gameOfLiberty, _ = ggol.NewGame(
		ggolSize,
		&initialUnit,
	)

	gsi.gameOfLiberty.SetNextUnitGenerator(gameNextUnitGenerator)

	gameField, _ := gsi.gameDao.GetGameField()
	gameUnits := convertGameFieldToGameUnits(gameField)

	for x := 0; x < ggolSize.Width; x += 1 {
		for y := 0; y < ggolSize.Height; y += 1 {
			gameFieldUnit := &(*gameUnits)[x][y]
			coord := &ggol.Coordinate{X: x, Y: y}
			gsi.gameOfLiberty.SetUnit(coord, gameFieldUnit)
		}
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

func (gsi *gameStoreImplement) GetGameFieldSize() *GameSize {
	return convertGgolSizeToGameSize(gsi.gameOfLiberty.GetSize())
}
