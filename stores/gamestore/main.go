package gamestore

import (
	"fmt"

	"github.com/DumDumGeniuss/game-of-liberty-computer/daos/gamedao"
	"github.com/DumDumGeniuss/ggol"
)

type GameStore interface {
	StartGame()
	StopGame()
	GetGameUnitsInArea()
	GetGameFieldSize()
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
	gameSize := convertGameFieldSizeToGameSize(gameFieldSize)

	gsi.gameOfLife, _ = ggol.NewGame(
		&ggol.Size{Width: gameSize.Width, Height: gameSize.Height},
		&initialUnit,
	)

	for i := 0; i < gameSize.Width; i += 1 {
		for j := 0; j < gameSize.Height; j += 1 {
			gsi.gameOfLife.SetUnit(&ggol.Coordinate{X: i, Y: j}, &(*gameUnits)[i][j])
		}
	}
}

func (gsi *gameStoreImplement) StartGame() {
	if gsi.gameOfLife == nil {
		gsi.initializeGame()
	}

	fmt.Println(gsi.gameOfLife.GetUnits())
}

func (gsi *gameStoreImplement) StopGame() {

}

func (gsi *gameStoreImplement) GetGameUnitsInArea() {

}

func (gsi *gameStoreImplement) GetGameFieldSize() {

}
