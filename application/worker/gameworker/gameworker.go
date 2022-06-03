package gameworker

import (
	"time"

	"github.com/DumDumGeniuss/game-of-liberty-computer/application/service/messageservice"
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/service/messageservicetopic"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/service/gameroomservice"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/config"
)

type GameWorker interface {
	StartGame() error
	StopGame() error
}

type gameWorkerImpl struct {
	gameRoomService gameroomservice.GameRoomService
	messageService  messageservice.MessageService
	gameTicker      *time.Ticker
	gameTickerStop  chan bool
}

var gameWorker GameWorker

func GetGameWorker(gameRoomService gameroomservice.GameRoomService, messageService messageservice.MessageService) GameWorker {
	if gameWorker == nil {
		gameWorker = &gameWorkerImpl{
			gameRoomService: gameRoomService,
			messageService:  messageService,
		}
	}
	return gameWorker
}

func (gwi *gameWorkerImpl) StartGame() error {
	gameId := config.GetConfig().GetGameId()

	go func() {
		gwi.gameTicker = time.NewTicker(time.Millisecond * 1000)
		defer gwi.gameTicker.Stop()
		gwi.gameTickerStop = make(chan bool)

		for {
			select {
			case <-gwi.gameTicker.C:
				gwi.gameRoomService.GenerateNextGameUnitMatrix(gameId)

				gwi.messageService.Publish(messageservicetopic.GameWorkerTickedMessageTopic, nil)
			case <-gwi.gameTickerStop:
				gwi.gameTicker.Stop()
				gwi.gameTicker = nil
				return
			}
		}
	}()

	return nil
}

func (gwi *gameWorkerImpl) StopGame() error {
	if gwi.gameTicker == nil {
		return nil
	}

	gwi.gameTickerStop <- true
	close(gwi.gameTickerStop)

	return nil
}
