package tickunitmapjob

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/infrastructure/memoryrepository"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/application/eventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/repository"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/memoryeventbus"
)

type Job interface {
	Start()
	Stop() error
}

type jobImplement struct {
	gameRoomRepository repository.GameRoomRepository
	eventBus           eventbus.EventBus
	gameTicker         *time.Ticker
	unitMapTickerStop  chan bool
}

var job *jobImplement

func GetJob() Job {
	if job == nil {
		gameRoomRepositoryMemory := memoryrepository.GetGameRoomRepositoryMemory()
		memoryEventBus := memoryeventbus.GetEventBus()

		job = &jobImplement{
			gameRoomRepository: gameRoomRepositoryMemory,
			eventBus:           memoryEventBus,
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
				gameRoomApplicationService := applicationservice.NewGameRoomApplicationService(
					applicationservice.GameRoomApplicationServiceConfiguration{
						GameRoomRepository: gwi.gameRoomRepository,
						EventBus:           gwi.eventBus,
					},
				)
				gameRoomApplicationService.TcikAllUnitMaps()
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
