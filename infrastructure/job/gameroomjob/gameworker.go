package gameroomjob

import (
	"time"

	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/service/gameroomservice"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/service/messageservice"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/service/messageservicetopic"
	"github.com/google/uuid"
)

type GameRoomJob interface {
	StartGame(gameId uuid.UUID) error
	StopGame(gameId uuid.UUID) error
}

type gameRoomJobImpl struct {
	gameRoomService gameroomservice.GameRoomService
	messageService  messageservice.MessageService
	gameTicker      map[uuid.UUID]*time.Ticker
	gameTickerStop  map[uuid.UUID]chan bool
}

var gameRoomJob GameRoomJob

func NewGameRoomJob(gameRoomService gameroomservice.GameRoomService, messageService messageservice.MessageService) GameRoomJob {
	if gameRoomJob == nil {
		gameRoomJob = &gameRoomJobImpl{
			gameRoomService: gameRoomService,
			messageService:  messageService,
			gameTicker:      make(map[uuid.UUID]*time.Ticker),
			gameTickerStop:  make(map[uuid.UUID]chan bool),
		}
	}
	return gameRoomJob
}

func (gwi *gameRoomJobImpl) StartGame(gameId uuid.UUID) error {
	go func() {
		gwi.gameTicker[gameId] = time.NewTicker(time.Millisecond * 1000)
		defer gwi.gameTicker[gameId].Stop()
		gwi.gameTickerStop[gameId] = make(chan bool)

		for {
			select {
			case <-gwi.gameTicker[gameId].C:
				gwi.gameRoomService.GenerateNextGameUnitMatrix(gameId)

				gwi.messageService.Publish(messageservicetopic.GameRoomJobTickedMessageTopic, nil)
			case <-gwi.gameTickerStop[gameId]:
				gwi.gameTicker[gameId].Stop()
				gwi.gameTicker = nil
				return
			}
		}
	}()

	return nil
}

func (gwi *gameRoomJobImpl) StopGame(gameId uuid.UUID) error {
	if gwi.gameTicker[gameId] == nil {
		return nil
	}

	gwi.gameTickerStop[gameId] <- true
	close(gwi.gameTickerStop[gameId])

	return nil
}
