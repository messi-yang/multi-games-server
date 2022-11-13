package livegameserver

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/port/dto"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegameserver/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegameserver/interface/applicationeventhandler"
)

func Start() {
	liveGameApplicationService, _ := applicationservice.NewLiveGameApplicationService(
		applicationservice.WithLiveGameService(),
	)

	liveGameId, _ := liveGameApplicationService.CreateLiveGame(dto.DimensionDto{
		Width:  200,
		Height: 200,
	})

	// gameService, _ := service.NewGameService(
	// 	service.WithPostgresGameRepository(),
	// )
	// games, _ := gameService.GeAllGames()
	// fmt.Println(games)

	applicationeventhandler.NewGameApplicationEventHandler(
		applicationeventhandler.GameApplicationEventHandlerConfiguration{
			LiveGameApplicationService: liveGameApplicationService,
		},
		liveGameId,
	)
}
