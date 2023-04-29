package gamesockethandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/gameappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/service/gamedomainsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/memrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/pguow"
)

func provideGameAppService(pgUow pguow.Uow) gameappsrv.Service {
	itemRepo := pgrepo.NewItemRepo(pgUow)
	playerRepo := memrepo.NewPlayerMemRepo()
	worldRepo := pgrepo.NewWorldRepo(pgUow)
	unitRepo := pgrepo.NewUnitRepo(pgUow)
	gameDomainService := gamedomainsrv.NewService(worldRepo, playerRepo, unitRepo, itemRepo)
	return gameappsrv.NewService(
		worldRepo, playerRepo, unitRepo, itemRepo, gameDomainService,
	)
}
