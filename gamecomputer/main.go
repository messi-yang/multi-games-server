package gamecomputer

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/game/port/dto"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/infrastructure/config"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/interface/applicationeventhandler"
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
