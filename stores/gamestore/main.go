package gamestore

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/daos/gamedao"
	"github.com/DumDumGeniuss/ggol"
)

type GameStore interface {
	StartGame()
	StopGame()
	GetGameUnitsInArea(area *GameArea) (*[][]*GameUnit, error)
	GetGameFieldSize() *GameSize
}

type gameStoreImplement struct {
	gameOfLife ggol.Game[GameUnit]
}

var Store GameStore = &gameStoreImplement{}

func (gsi *gameStoreImplement) initializeGame() {
	gameField, _ := gamedao.DAO.GetGameField()
	gameFieldSize, _ := gamedao.DAO.GetGameFieldSize()

	initialUnit := GameUnit{
		Alive: false,
		Age:   0,
	}
	gameUnits := convertGameFieldToGameUnits(gameField)
	ggolSize := convertGameGameFieldSizeToGgolSize(gameFieldSize)

	gsi.gameOfLife, _ = ggol.NewGame(
		ggolSize,
		&initialUnit,
	)

	for x := 0; x < ggolSize.Width; x += 1 {
		for y := 0; y < ggolSize.Height; y += 1 {
			gameFieldUnit := &(*gameUnits)[x][y]
			coord := &ggol.Coordinate{X: x, Y: y}
			gsi.gameOfLife.SetUnit(coord, gameFieldUnit)
		}
	}
}

func (gsi *gameStoreImplement) StartGame() {
	if gsi.gameOfLife == nil {
		gsi.initializeGame()
	}

	// fmt.Println(gsi.gameOfLife.GetUnits())
}

func (gsi *gameStoreImplement) StopGame() {

}

func (gsi *gameStoreImplement) GetGameUnitsInArea(area *GameArea) (*[][]*GameUnit, error) {
	ggolArea := convertGameAreaToGgolArea(area)
	units, err := gsi.gameOfLife.GetUnitsInArea(ggolArea)
	if err != nil {
		return nil, err
	}
	return &units, nil
}

func (gsi *gameStoreImplement) GetGameFieldSize() *GameSize {
	return convertGgolSizeToGameSize(gsi.gameOfLife.GetSize())
}
