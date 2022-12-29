package gameserver

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/infrastructure/integrationevent/redisintegrationeventpublisher"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/infrastructure/persistence/gamepsqlrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/service/gamedomainservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/gameserver/application/service/livegameappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/gameserver/infrastructure/persistence/livegamememoryrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/gameserver/interface/livegameeventcontroller"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/gormdb"
)

func Start() {
	gormDb, err := gormdb.New()
	if err != nil {
		panic(err)
	}
	gameRepo := gamepsqlrepo.New(gormDb)
	liveGameRepo := livegamememoryrepo.New()
	gameDomainService := gamedomainservice.New(
		gameRepo,
	)
	notificationPublisher := redisintegrationeventpublisher.New()
	liveGameAppService := livegameappservice.New(
		liveGameRepo,
		gameRepo,
		notificationPublisher,
	)

	games, _ := gameRepo.GetAll()
	if len(games) > 0 {
		liveGameAppService.CreateLiveGame(games[0].GetId())
	} else {
		dimension, _ := commonmodel.NewDimension(200, 200)
		gameId, _ := gameDomainService.CreateGame(dimension)
		livegamemodel.NewLiveGameId(gameId.GetId().String())
	}

	livegameeventcontroller.New(liveGameAppService)
}
