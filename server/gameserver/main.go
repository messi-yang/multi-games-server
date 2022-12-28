package gameserver

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/domainmodel/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/domainmodel/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/domainservice"
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/server/common/adapter/notification/redis"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/common/adapter/persistence/postgres"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/gameserver/adapter/persistence/memory"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/gameserver/application/appservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/gameserver/interface/eventcontroller"
)

func Start() {
	postgresClient, err := postgres.NewPostgresClient()
	if err != nil {
		panic(err)
	}
	gameRepository := postgres.NewPostgresGameRepository(postgresClient)
	liveGameRepository := memory.NewMemoryLiveGameRepository()
	gameDomainService := domainservice.NewGameDomainService(
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
		dimension, _ := commonmodel.NewDimension(200, 200)
		gameId, _ := gameDomainService.CreateGame(dimension)
		livegamemodel.NewLiveGameId(gameId.GetId().String())
	}

	eventcontroller.NewLiveGameEventController(liveGameAppService)
}
