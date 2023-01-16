package gameserver

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/infrastructure/messaging/redisintgreventpublisher"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/infrastructure/persistence/gamepsqlrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/gameserver/application/service/livegameappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/gameserver/infrastructure/persistence/livegamememoryrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/gameserver/interface/livegameintgreventcontroller"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/gormdb"
)

func Start() {
	gormDb, err := gormdb.New()
	if err != nil {
		panic(err)
	}
	gameRepo := gamepsqlrepo.New(gormDb)
	liveGameRepo := livegamememoryrepo.New()
	gameDomainService := service.NewGameDomainService(
		gameRepo,
	)
	intgrEventPublisher := redisintgreventpublisher.New()
	liveGameAppService := livegameappservice.New(
		liveGameRepo,
		gameRepo,
		intgrEventPublisher,
	)

	games, _ := gameRepo.GetAll()
	if len(games) > 0 {
		liveGameAppService.CreateLiveGame(games[0].GetId().ToString())
	} else {
		mapSize, _ := commonmodel.NewSizeVo(200, 200)
		gameId, _ := gameDomainService.CreateGame(mapSize)
		livegamemodel.NewLiveGameIdVo(gameId.ToString())
	}

	livegameintgreventcontroller.New(liveGameAppService)
}
