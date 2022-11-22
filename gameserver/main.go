package gameserver

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamecommonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/service/gameservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/gameserver/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/gameserver/interface/applicationeventhandler"
)

func Start() {
	liveGameApplicationService, _ := applicationservice.NewLiveGameApplicationService(
		applicationservice.WithLiveGameService(),
		applicationservice.WithGameService(),
		applicationservice.WithRedisNotificationPublisher(),
	)

	gameService, _ := gameservice.NewGameService(
		gameservice.WithPostgresGameRepository(),
	)
	games, _ := gameService.GeAllGames()
	var liveGameId livegamemodel.LiveGameId
	if len(games) > 0 {
		liveGameId, _ = liveGameApplicationService.CreateLiveGame(games[0].GetId())
	} else {
		dimension, _ := gamecommonmodel.NewDimension(200, 200)
		gameId, _ := gameService.CreateGame(dimension)
		liveGameId = livegamemodel.NewLiveGameId(gameId.GetId())
	}

	applicationeventhandler.NewGameIntegrationEventHandler(
		applicationeventhandler.GameIntegrationEventHandlerConfiguration{
			LiveGameApplicationService: liveGameApplicationService,
		},
		liveGameId,
	)
}
