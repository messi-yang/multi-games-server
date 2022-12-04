package gameserver

import (
	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/service/gameservice"
	applicationservice "github.com/dum-dum-genius/game-of-liberty-computer/server/gameserver/application/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/gameserver/interface/eventcontroller"
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

	eventcontroller.NewLiveGameEventController(
		eventcontroller.LiveGameEventControllerConfiguration{
			LiveGameApplicationService: liveGameApplicationService,
		},
	)
}
