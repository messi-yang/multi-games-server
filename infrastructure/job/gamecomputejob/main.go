package gamecomputejob

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/application/event/gamecomputedevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/usecase/tickallunitmapssusecase"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/repository/gameroomrepository"
	"github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/eventbus/gamecomputedeventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/memory/gameroommemory"
)

type GameComputeJob interface {
	Start()
	Stop() error
}

type gameComputeJobImpl struct {
	gameRoomRepository  gameroomrepository.GameRoomRepository
	gameComputeEventBus gamecomputedevent.GameComputedEvent
	gameTicker          *time.Ticker
	gameTickerStop      chan bool
}

var gameComputeJob *gameComputeJobImpl

func GetGameComputeJob() GameComputeJob {
	if gameComputeJob == nil {
		gameRoomMemory := gameroommemory.GetGameRoomMemory()
		gameComputeEventBus := gamecomputedeventbus.GetGameComputedEventBus()

		gameComputeJob = &gameComputeJobImpl{
			gameRoomRepository:  gameRoomMemory,
			gameComputeEventBus: gameComputeEventBus,
		}

		return gameComputeJob
	}

	return gameComputeJob
}

func (gwi *gameComputeJobImpl) Start() {
	go func() {
		gwi.gameTicker = time.NewTicker(time.Millisecond * 1000)
		defer gwi.gameTicker.Stop()
		gwi.gameTickerStop = make(chan bool)

		for {
			select {
			case <-gwi.gameTicker.C:
				computeAllGamesUseCase := tickallunitmapssusecase.New(gwi.gameRoomRepository, gwi.gameComputeEventBus)
				computeAllGamesUseCase.Execute()
			case <-gwi.gameTickerStop:
				gwi.gameTicker.Stop()
				gwi.gameTicker = nil
			}
		}
	}()
}

func (gwi *gameComputeJobImpl) Stop() error {
	if gwi.gameTicker == nil {
		return nil
	}

	gwi.gameTickerStop <- true
	close(gwi.gameTickerStop)

	return nil
}
