package gamesockethandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/gameappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/service/gamedomainsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/postgres/pgrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/event/memory/memdomainevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
)

func provideGameAppService(uow pguow.Uow) gameappsrv.Service {
	itemRepo := pgrepo.NewItemRepo(uow)
	playerRepo := pgrepo.NewPlayerRepo(uow)
	worldRepo := pgrepo.NewWorldRepo(uow)
	unitRepo := pgrepo.NewUnitRepo(uow)
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	gameDomainService := gamedomainsrv.NewService(worldRepo, playerRepo, unitRepo, itemRepo, domainEventDispatcher)
	return gameappsrv.NewService(
		worldRepo, playerRepo, unitRepo, itemRepo, gameDomainService,
	)
}
