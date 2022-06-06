package gameupdatejob

import (
	"time"

	"github.com/DumDumGeniuss/game-of-liberty-computer/application/event/gameupdateevent"
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/usecase/updateallgamesusecase"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/service/gameroomservice"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/memory/gameroommemory"
)

type GameUpdateJob interface {
	Start()
	Stop() error
}

type gameUpdateJobImpl struct {
	gameRoomService gameroomservice.GameRoomService
	eventBus        *gameupdateevent.GameUpdateEvent
	gameTicker      *time.Ticker
	gameTickerStop  chan bool
}

var gameUpdateJob *gameUpdateJobImpl

func GetGameUpdateJob() GameUpdateJob {
	if gameUpdateJob == nil {
		gameRoomMemoryRepository := gameroommemory.GetGameRoomMemoryRepository()
		gameRoomService := gameroomservice.NewGameRoomService(gameRoomMemoryRepository)
		eventBus := gameupdateevent.GetGameUpdateEventBus()

		gameUpdateJob = &gameUpdateJobImpl{
			gameRoomService: gameRoomService,
			eventBus:        eventBus,
		}

		return gameUpdateJob
	}

	return gameUpdateJob
}

func (gwi *gameUpdateJobImpl) Start() {
	go func() {
		gwi.gameTicker = time.NewTicker(time.Millisecond * 1000)
		defer gwi.gameTicker.Stop()
		gwi.gameTickerStop = make(chan bool)

		for {
			select {
			case <-gwi.gameTicker.C:
				updateAllGamesUseCase := updateallgamesusecase.NewUseCase(gwi.gameRoomService, gwi.eventBus)
				updateAllGamesUseCase.Execute()
			case <-gwi.gameTickerStop:
				gwi.gameTicker.Stop()
				gwi.gameTicker = nil
			}
		}
	}()
}

func (gwi *gameUpdateJobImpl) Stop() error {
	if gwi.gameTicker == nil {
		return nil
	}

	gwi.gameTickerStop <- true
	close(gwi.gameTickerStop)

	return nil
}
