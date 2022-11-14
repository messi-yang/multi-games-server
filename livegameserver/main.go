package livegameserver

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/common/domain/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/service/gameservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegameserver/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegameserver/interface/applicationeventhandler"
)

func Start() {
	liveGameApplicationService, _ := applicationservice.NewLiveGameApplicationService(
		applicationservice.WithLiveGameService(),
		applicationservice.WithGameService(),
	)

	gameService, _ := gameservice.NewGameService(
		gameservice.WithPostgresGameRepository(),
	)
	games, _ := gameService.GeAllGames()
	var liveGameId livegamemodel.LiveGameId
	if len(games) > 0 {
		liveGameId, _ = liveGameApplicationService.CreateLiveGame(games[0].GetId())
	} else {
		dimension, _ := valueobject.NewDimension(200, 200)
		gameId, _ := gameService.CreateGame(dimension)
		liveGameId = livegamemodel.NewLiveGameId(gameId.GetId())
	}

	applicationeventhandler.NewGameApplicationEventHandler(
		applicationeventhandler.GameApplicationEventHandlerConfiguration{
			LiveGameApplicationService: liveGameApplicationService,
		},
		liveGameId,
	)
}
