package tickunitmapjob

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/application/event/gameunitmaptickedevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/service/gameroomservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/repository/gameroomrepository"
	"github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/eventbus/gameunitmaptickedeventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/memory/gameroommemory"
)

type TickUnitMapJob interface {
	Start()
	Stop() error
}

type tickUnitMapJobImpl struct {
	gameRoomRepository        gameroomrepository.Repository
	gameUnitMapTickedEventBus gameunitmaptickedevent.Event
	gameTicker                *time.Ticker
	unitMapTickerStop         chan bool
}

var tickUnitMapJob *tickUnitMapJobImpl

func GetTickUnitMapJob() TickUnitMapJob {
	if tickUnitMapJob == nil {
		gameRoomMemory := gameroommemory.GetGameRoomMemory()
		gameUnitMapTickedEventBus := gameunitmaptickedeventbus.GetGameUnitMapTickedEventBus()

		tickUnitMapJob = &tickUnitMapJobImpl{
			gameRoomRepository:        gameRoomMemory,
			gameUnitMapTickedEventBus: gameUnitMapTickedEventBus,
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
				gameRoomService := gameroomservice.NewService(
					gameroomservice.Configuration{
						GameRoomRepository:     gwi.gameRoomRepository,
						GameUnitMapTickedEvent: gwi.gameUnitMapTickedEventBus,
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
