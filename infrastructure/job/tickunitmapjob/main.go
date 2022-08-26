package tickunitmapjob

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/application/event/gameunitmapupdatedevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/service/gameroomservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/repository/gameroomrepository"
	"github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/eventbus/gameunitmapupdatedeventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/memory/gameroommemory"
)

type TickUnitMapJob interface {
	Start()
	Stop() error
}

type tickUnitMapJobImpl struct {
	gameRoomRepository         gameroomrepository.GameRoomRepository
	gameUnitMapUpdatedEventBus gameunitmapupdatedevent.GameUnitMapUpdatedEvent
	gameTicker                 *time.Ticker
	unitMapTickerStop          chan bool
}

var tickUnitMapJob *tickUnitMapJobImpl

func GetTickUnitMapJob() TickUnitMapJob {
	if tickUnitMapJob == nil {
		gameRoomMemory := gameroommemory.GetGameRoomMemory()
		gameUnitMapUpdatedEventBus := gameunitmapupdatedeventbus.GetGameUnitMapUpdatedEventBus()

		tickUnitMapJob = &tickUnitMapJobImpl{
			gameRoomRepository:         gameRoomMemory,
			gameUnitMapUpdatedEventBus: gameUnitMapUpdatedEventBus,
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
				gameRoomService := gameroomservice.NewGameRoomService(
					gameroomservice.Configuration{
						GameRoomRepository: gwi.gameRoomRepository,
						GameComputeEvent:   gwi.gameUnitMapUpdatedEventBus,
					},
				)
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
