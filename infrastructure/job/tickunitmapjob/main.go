package tickunitmapjob

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/application/event/gameunitmaptickedevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/service/gameroomservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/repository/gameroomrepository"
	"github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/eventbus/gameunitmaptickedeventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/memory/gameroommemory"
)

type Job interface {
	Start()
	Stop() error
}

type jobImplement struct {
	gameRoomRepository        gameroomrepository.Repository
	gameUnitMapTickedEventBus gameunitmaptickedevent.Event
	gameTicker                *time.Ticker
	unitMapTickerStop         chan bool
}

var job *jobImplement

func GetJob() Job {
	if job == nil {
		gameRoomRepository := gameroommemory.GetRepository()
		gameUnitMapTickedEventBus := gameunitmaptickedeventbus.GetGameUnitMapTickedEventBus()

		job = &jobImplement{
			gameRoomRepository:        gameRoomRepository,
			gameUnitMapTickedEventBus: gameUnitMapTickedEventBus,
		}

		return job
	}

	return job
}

func (gwi *jobImplement) Start() {
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

func (gwi *jobImplement) Stop() error {
	if gwi.gameTicker == nil {
		return nil
	}

	gwi.unitMapTickerStop <- true
	close(gwi.unitMapTickerStop)

	return nil
}
