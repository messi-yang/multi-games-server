package livegameserver

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/port/dto"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegameserver/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegameserver/infrastructure/config"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegameserver/interface/applicationeventhandler"
)

func Start() {
	liveGameApplicationService, _ := applicationservice.NewLiveGameApplicationService(
		applicationservice.WithLiveGameService(),
	)

	size := config.GetConfig().GetLiveGameDimension()
	gameId, _ := liveGameApplicationService.CreateLiveGame(dto.DimensionDto{
		Width:  size,
		Height: size,
	})

	applicationeventhandler.NewGameApplicationEventHandler(
		applicationeventhandler.GameApplicationEventHandlerConfiguration{
			LiveGameApplicationService: liveGameApplicationService,
		},
		gameId,
	)
}
