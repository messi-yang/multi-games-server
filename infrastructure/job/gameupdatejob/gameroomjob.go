package gameupdatejob

import (
	"time"

	"github.com/DumDumGeniuss/game-of-liberty-computer/application/event/gameupdateevent"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/service/gameroomservice"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/service/messageservice"
)

type GameUpdateJob interface {
	Start() error
	Stop() error
}

type gameUpdateJobImpl struct {
	gameRoomService gameroomservice.GameRoomService
	messageService  messageservice.MessageService
	gameTicker      *time.Ticker
	gameTickerStop  chan bool
}

var gameUpdateJob GameUpdateJob

func NewGameUpdateJob(gameRoomService gameroomservice.GameRoomService, messageService messageservice.MessageService) GameUpdateJob {
	if gameUpdateJob == nil {
		gameUpdateJob = &gameUpdateJobImpl{
			gameRoomService: gameRoomService,
			messageService:  messageService,
		}
	}
	return gameUpdateJob
}

func (gwi *gameUpdateJobImpl) Start() error {
	eventBus := gameupdateevent.GetGameUpdateEventBus()
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
					eventBus.Publish(gameRoom.GetGameId())
				}
			case <-gwi.gameTickerStop:
				gwi.gameTicker.Stop()
				gwi.gameTicker = nil
				return
			}
		}
	}()

	return nil
}

func (gwi *gameUpdateJobImpl) Stop() error {
	if gwi.gameTicker == nil {
		return nil
	}

	gwi.gameTickerStop <- true
	close(gwi.gameTickerStop)

	return nil
}
