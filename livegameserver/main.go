package livegameserver

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/common/domain/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/game/domain/service"
	liveGameValueObject "github.com/dum-dum-genius/game-of-liberty-computer/livegame/domain/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegameserver/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegameserver/interface/applicationeventhandler"
)

func Start() {
	liveGameApplicationService, _ := applicationservice.NewLiveGameApplicationService(
		applicationservice.WithLiveGameService(),
		applicationservice.WithGameService(),
	)

	gameService, _ := service.NewGameService(
		service.WithPostgresGameRepository(),
	)
	games, _ := gameService.GeAllGames()
	var liveGameId liveGameValueObject.LiveGameId
	if len(games) > 0 {
		liveGameId, _ = liveGameApplicationService.CreateLiveGame(games[0].GetId())
	} else {
		dimension, _ := valueobject.NewDimension(200, 200)
		gameId, _ := gameService.CreateGame(dimension)
		liveGameId = liveGameValueObject.NewLiveGameId(gameId.GetId())
	}

	applicationeventhandler.NewGameApplicationEventHandler(
		applicationeventhandler.GameApplicationEventHandlerConfiguration{
			LiveGameApplicationService: liveGameApplicationService,
		},
		liveGameId,
	)
}
