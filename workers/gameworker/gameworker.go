package gameworker

import (
	"fmt"
	"time"

	"github.com/DumDumGeniuss/game-of-liberty-computer/providers/gameprovider"
)

type GameWorker interface {
	SetGameStore(gameprovider.GameStore)
	Start() error
	Stop() error
}

type gameWorkerImpl struct {
	gameStore      gameprovider.GameStore
	gameTicker     *time.Ticker
	gameTickerStop chan bool
}

var Worker GameWorker = &gameWorkerImpl{
	gameStore:      nil,
	gameTicker:     nil,
	gameTickerStop: nil,
}

func (gwi *gameWorkerImpl) SetGameStore(gameStore gameprovider.GameStore) {
	gwi.gameStore = gameStore
}

func (gwi *gameWorkerImpl) Start() error {
	if gwi.gameStore == nil {
		return &errGameStoreIsNotSet{}
	}
	go func() {
		gwi.gameTicker = time.NewTicker(time.Millisecond * 1000)
		defer gwi.gameTicker.Stop()
		gwi.gameTickerStop = make(chan bool)

		for {
			select {
			case <-gwi.gameTicker.C:
				fmt.Println("hi")
			case <-gwi.gameTickerStop:
				gwi.gameTicker.Stop()
				gwi.gameTicker = nil
				return
			}
		}
	}()

	return nil
}

func (gwi *gameWorkerImpl) Stop() error {
	if gwi.gameStore == nil {
		return &errGameStoreIsNotSet{}
	}
	if gwi.gameTicker == nil {
		return nil
	}
	gwi.gameTickerStop <- true
	close(gwi.gameTickerStop)

	return nil
}
