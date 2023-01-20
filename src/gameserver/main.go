package main

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/infrastructure/psqlrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/infrastructure/redispub"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/gameserver/application/appservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/gameserver/infrastructure/memrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/gameserver/interface/inteventcontroller"
)

func main() {
	gameRepo, _ := psqlrepo.NewGamePsqlRepo()
	liveGameRepo := memrepo.NewLiveGameMemRepo()
	gameDomainService := service.NewGameService(
		gameRepo,
	)
	IntEventPublisher := redispub.New()
	liveGameAppService := appservice.NewLiveGameAppService(
		liveGameRepo,
		gameRepo,
		IntEventPublisher,
	)

	games, _ := gameRepo.GetAll()
	if len(games) > 0 {
		liveGameAppService.CreateLiveGame(games[0].GetId().ToString())
	} else {
		mapSize, _ := commonmodel.NewSizeVo(200, 200)
		gameId, _ := gameDomainService.CreateGame(mapSize)
		livegamemodel.NewLiveGameIdVo(gameId.ToString())
	}

	inteventcontroller.NewLiveGameIntEventController(liveGameAppService)
}
