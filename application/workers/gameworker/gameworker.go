package gameworker

import (
	"time"

	"github.com/DumDumGeniuss/game-of-liberty-computer/application/services/messageservice"
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/services/messageservicetopic"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/service/gameroomservice"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/config"
)

type GameWorker interface {
	InjectGameRoomService(gameRoomService gameroomservice.GameRoomService)
	InjectMessageService(messageService messageservice.MessageService)
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

func GetGameWorker() GameWorker {
	if gameWorker == nil {
		gameWorker = &gameWorkerImpl{}
	}
	return gameWorker
}

func (gwi *gameWorkerImpl) InjectGameRoomService(gameRoomService gameroomservice.GameRoomService) {
	gwi.gameRoomService = gameRoomService
}
func (gwi *gameWorkerImpl) InjectMessageService(messageService messageservice.MessageService) {
	gwi.messageService = messageService
}

func (gwi *gameWorkerImpl) checkGameServiceDependency() error {
	if gwi.gameRoomService == nil {
		return &errMissingGameServiceDependency{}
	}

	return nil
}

func (gwi *gameWorkerImpl) checkMessageServiceDependency() error {
	if gwi.messageService == nil {
		return &errMissingMessageServiceDependency{}
	}

	return nil
}

func (gwi *gameWorkerImpl) StartGame() error {
	if err := gwi.checkGameServiceDependency(); err != nil {
		return err
	}
	if err := gwi.checkMessageServiceDependency(); err != nil {
		return err
	}

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
