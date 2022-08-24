package tickunitmapjob

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/application/event/gamecomputedevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/service/gameroomservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/repository/gameroomrepository"
	"github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/eventbus/gamecomputedeventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/memory/gameroommemory"
)

type TickUnitMapJob interface {
	Start()
	Stop() error
}

type tickUnitMapJobImpl struct {
	gameRoomRepository  gameroomrepository.GameRoomRepository
	gameComputeEventBus gamecomputedevent.GameComputedEvent
	gameTicker          *time.Ticker
	unitMapTickerStop   chan bool
}

var tickUnitMapJob *tickUnitMapJobImpl

func GetTickUnitMapJob() TickUnitMapJob {
	if tickUnitMapJob == nil {
		gameRoomMemory := gameroommemory.GetGameRoomMemory()
		gameComputeEventBus := gamecomputedeventbus.GetGameComputedEventBus()

		tickUnitMapJob = &tickUnitMapJobImpl{
			gameRoomRepository:  gameRoomMemory,
			gameComputeEventBus: gameComputeEventBus,
		}

		return tickUnitMapJob
	}

	return tickUnitMapJob
}

func (gwi *tickUnitMapJobImpl) Start() {
	go func() {
		gwi.gameTicker = time.NewTicker(time.Millisecond * 1000)
		defer gwi.gameTicker.Stop()
		gwi.unitMapTickerStop = make(chan bool)

		for {
			select {
			case <-gwi.gameTicker.C:
				gameRoomService := gameroomservice.NewGameRoomServiceWithGameComputedEvent(gwi.gameRoomRepository, gwi.gameComputeEventBus)
				gameRoomService.TcikAllUnitMaps()
			case <-gwi.unitMapTickerStop:
				gwi.gameTicker.Stop()
				gwi.gameTicker = nil
			}
		}
	}()
}

func (gwi *tickUnitMapJobImpl) Stop() error {
	if gwi.gameTicker == nil {
		return nil
	}

	gwi.unitMapTickerStop <- true
	close(gwi.unitMapTickerStop)

	return nil
}
