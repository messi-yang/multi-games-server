package gameworker

import (
	"time"

	"github.com/DumDumGeniuss/game-of-liberty-computer/application/service/messageservice"
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/service/messageservicetopic"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/service/gameroomservice"
	"github.com/google/uuid"
)

type GameWorker interface {
	StartGame(gameId uuid.UUID) error
	StopGame(gameId uuid.UUID) error
}

type gameWorkerImpl struct {
	gameRoomService gameroomservice.GameRoomService
	messageService  messageservice.MessageService
	gameTicker      map[uuid.UUID]*time.Ticker
	gameTickerStop  map[uuid.UUID]chan bool
}

var gameWorker GameWorker

func GetGameWorker(gameRoomService gameroomservice.GameRoomService, messageService messageservice.MessageService) GameWorker {
	if gameWorker == nil {
		gameWorker = &gameWorkerImpl{
			gameRoomService: gameRoomService,
			messageService:  messageService,
			gameTicker:      make(map[uuid.UUID]*time.Ticker),
			gameTickerStop:  make(map[uuid.UUID]chan bool),
		}
	}
	return gameWorker
}

func (gwi *gameWorkerImpl) StartGame(gameId uuid.UUID) error {
	go func() {
		gwi.gameTicker[gameId] = time.NewTicker(time.Millisecond * 1000)
		defer gwi.gameTicker[gameId].Stop()
		gwi.gameTickerStop[gameId] = make(chan bool)

		for {
			select {
			case <-gwi.gameTicker[gameId].C:
				gwi.gameRoomService.GenerateNextGameUnitMatrix(gameId)

				gwi.messageService.Publish(messageservicetopic.GameWorkerTickedMessageTopic, nil)
			case <-gwi.gameTickerStop[gameId]:
				gwi.gameTicker[gameId].Stop()
				gwi.gameTicker = nil
				return
			}
		}
	}()

	return nil
}

func (gwi *gameWorkerImpl) StopGame(gameId uuid.UUID) error {
	if gwi.gameTicker[gameId] == nil {
		return nil
	}

	gwi.gameTickerStop[gameId] <- true
	close(gwi.gameTickerStop[gameId])

	return nil
}
