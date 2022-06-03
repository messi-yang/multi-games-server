package gameroomjob

import (
	"time"

	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/service/gameroomservice"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/service/messageservice"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/service/messageservicetopic"
)

type GameRoomJob interface {
	Start() error
	Stop() error
}

type gameRoomJobImpl struct {
	gameRoomService gameroomservice.GameRoomService
	messageService  messageservice.MessageService
	gameTicker      *time.Ticker
	gameTickerStop  chan bool
}

var gameRoomJob GameRoomJob

func NewGameRoomJob(gameRoomService gameroomservice.GameRoomService, messageService messageservice.MessageService) GameRoomJob {
	if gameRoomJob == nil {
		gameRoomJob = &gameRoomJobImpl{
			gameRoomService: gameRoomService,
			messageService:  messageService,
		}
	}
	return gameRoomJob
}

func (gwi *gameRoomJobImpl) Start() error {
	go func() {
		gwi.gameTicker = time.NewTicker(time.Millisecond * 1000)
		defer gwi.gameTicker.Stop()
		gwi.gameTickerStop = make(chan bool)

		for {
			select {
			case <-gwi.gameTicker.C:
				gameRooms := gwi.gameRoomService.GetAllGameRooms()
				for _, gameRoom := range gameRooms {
					gwi.gameRoomService.GenerateNextGameUnitMatrix(gameRoom.GetGameId())
				}

				gwi.messageService.Publish(messageservicetopic.GameRoomJobTickedMessageTopic, nil)
			case <-gwi.gameTickerStop:
				gwi.gameTicker.Stop()
				gwi.gameTicker = nil
				return
			}
		}
	}()

	return nil
}

func (gwi *gameRoomJobImpl) Stop() error {
	if gwi.gameTicker == nil {
		return nil
	}

	gwi.gameTickerStop <- true
	close(gwi.gameTickerStop)

	return nil
}
