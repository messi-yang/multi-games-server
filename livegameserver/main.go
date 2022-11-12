package livegameserver

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/game/port/dto"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegameserver/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegameserver/infrastructure/config"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegameserver/interface/applicationeventhandler"
)

func Start() {
	gameApplicationService, _ := applicationservice.NewGameApplicationService(
		applicationservice.WithGameService(),
	)

	size := config.GetConfig().GetGameDimension()
	gameId, _ := gameApplicationService.CreateGame(dto.DimensionDto{
		Width:  size,
		Height: size,
	})

	applicationeventhandler.NewGameApplicationEventHandler(
		applicationeventhandler.GameApplicationEventHandlerConfiguration{
			GameApplicationService: gameApplicationService,
		},
		gameId,
	)
}
