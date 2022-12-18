package gameserver

import (
	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/service/gameservice"
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/persistence/postgres"
	appservice "github.com/dum-dum-genius/game-of-liberty-computer/server/gameserver/application/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/gameserver/interface/eventcontroller"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/gameserver/port/adapter/persistence/memory"
)

func Start() {
	postgresClient, err := postgres.NewPostgresClient()
	if err != nil {
		panic(err)
	}
	gameRepository := postgres.NewPostgresGameRepository(postgresClient)
	liveGameRepository := memory.NewMemoryLiveGameRepository()
	gameService := gameservice.NewGameService(
		gameRepository,
	)
	notificationPublisher := commonredis.NewRedisNotificationPublisher()
	liveGameAppService := appservice.NewLiveGameAppService(
		liveGameRepository,
		gameRepository,
		notificationPublisher,
	)

	games, _ := gameRepository.GetAll()
	if len(games) > 0 {
		liveGameAppService.CreateLiveGame(games[0].GetId())
	} else {
		dimension, _ := gamecommonmodel.NewDimension(200, 200)
		gameId, _ := gameService.CreateGame(dimension)
		livegamemodel.NewLiveGameId(gameId.GetId().String())
	}

	eventcontroller.NewLiveGameEventController(
		eventcontroller.LiveGameEventControllerConfiguration{
			LiveGameAppService: liveGameAppService,
		},
	)
}
