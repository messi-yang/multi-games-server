package gamestore

import (
	"time"

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
	gameOfLife     ggol.Game[GameUnit]
	gameDao        gamedao.GameDAO
	gameTicker     *time.Ticker
	gameTickerStop chan bool
}

var Store GameStore = &gameStoreImplement{
	gameOfLife:     nil,
	gameDao:        gamedao.DAO,
	gameTicker:     nil,
	gameTickerStop: nil,
}

func (gsi *gameStoreImplement) initializeGame() {
	initialUnit := GameUnit{
		Alive: false,
		Age:   0,
	}
	gameUnits := gsi.getGameUnitsFromGameDao()
	ggolSize := gsi.getGgolSizeFromGameDao()

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

func (gsi *gameStoreImplement) getGameUnitsFromGameDao() *GameUnits {
	gameField, _ := gsi.gameDao.GetGameField()
	gameUnits := convertGameFieldToGameUnits(gameField)

	return gameUnits
}

func (gsi *gameStoreImplement) getGgolSizeFromGameDao() *ggol.Size {
	gameFieldSize, _ := gsi.gameDao.GetGameFieldSize()
	ggolSize := convertGameGameFieldSizeToGgolSize(gameFieldSize)

	return ggolSize
}

func (gsi *gameStoreImplement) StartGame() {
	if gsi.gameOfLife == nil {
		gsi.initializeGame()
	}

	go func() {
		gsi.gameTicker = time.NewTicker(time.Millisecond * 1000)
		defer gsi.gameTicker.Stop()
		gsi.gameTickerStop = make(chan bool)

		for {
			select {
			case <-gsi.gameTicker.C:
				gsi.gameOfLife.GenerateNextUnits()
			case <-gsi.gameTickerStop:
				gsi.gameTicker.Stop()
				gsi.gameTicker = nil
				return
			}
		}
	}()
}

func (gsi *gameStoreImplement) StopGame() {
	if gsi.gameTicker == nil {
		return
	}
	gsi.gameTickerStop <- true
	close(gsi.gameTickerStop)
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
