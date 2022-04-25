package gameworker

import (
	"time"

	"github.com/DumDumGeniuss/game-of-liberty-computer/services/gameservice"
)

type GameWorker interface {
	StartGame()
	StopGame()
}

type gameWorkerImpl struct {
	gameService    gameservice.GameService
	gameTicker     *time.Ticker
	gameTickerStop chan bool
}

var gameWorker GameWorker

func CreateGameWorker(gameService gameservice.GameService) (GameWorker, error) {
	if gameWorker != nil {
		return nil, &errGameWorkHasBeenCreated{}
	}

	return &gameWorkerImpl{
		gameService: gameService,
	}, nil
}

func (gwi *gameWorkerImpl) StartGame() {
	go func() {
		gwi.gameTicker = time.NewTicker(time.Millisecond * 1000)
		defer gwi.gameTicker.Stop()
		gwi.gameTickerStop = make(chan bool)

		for {
			select {
			case <-gwi.gameTicker.C:
				gwi.gameService.GenerateNextUnits()

				// gameArea := gameservice.GameArea{
				// 	From: gameservice.GameCoordinate{X: 0, Y: 0},
				// 	To:   gameservice.GameCoordinate{X: 3, Y: 3},
				// }
				// gameUnits, err := gwi.gameService.GetGameUnitsInArea(&gameArea)
				// if err != nil {
				// 	fmt.Println(err.Error())
				// }
				// fmt.Println(gameUnits)
			case <-gwi.gameTickerStop:
				gwi.gameTicker.Stop()
				gwi.gameTicker = nil
				return
			}
		}
	}()
}

func (gwi *gameWorkerImpl) StopGame() {
	if gwi.gameTicker == nil {
		return
	}
	gwi.gameTickerStop <- true
	close(gwi.gameTickerStop)
}
