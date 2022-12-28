package gameserver

import (
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/src/common/adapter/notification/redis"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/infrastructure/persistence/postgres"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/service/gamedomainservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/gameserver/application/service/livegameappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/gameserver/infrastructure/persistence/livegamememoryrepository"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/gameserver/interface/livegameeventcontroller"
)

func Start() {
	postgresClient, err := postgres.NewPostgresClient()
	if err != nil {
		panic(err)
	}
	gameRepository := postgres.NewPostgresGameRepository(postgresClient)
	liveGameRepository := livegamememoryrepository.New()
	gameDomainService := gamedomainservice.New(
		gameRepository,
	)
	notificationPublisher := commonredis.NewRedisNotificationPublisher()
	liveGameAppService := livegameappservice.New(
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

	livegameeventcontroller.New(liveGameAppService)
}
