package gameserver

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/module/game/domain/model/gamecommonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/module/game/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/module/game/domain/service/gameservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/gameserver/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/gameserver/interface/applicationeventhandler"
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
	if len(games) > 0 {
		liveGameApplicationService.CreateLiveGame(games[0].GetId())
	} else {
		dimension, _ := gamecommonmodel.NewDimension(200, 200)
		gameId, _ := gameService.CreateGame(dimension)
		livegamemodel.NewLiveGameId(gameId.GetId())
	}

	applicationeventhandler.NewGameIntegrationEventHandler(
		applicationeventhandler.GameIntegrationEventHandlerConfiguration{
			LiveGameApplicationService: liveGameApplicationService,
		},
	)
}
