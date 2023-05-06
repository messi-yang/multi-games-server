package worldhttphandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/gamerappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/worldappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/messaging/memdomainevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/pguow"
)

func provideGamerAppService(uow pguow.Uow) gamerappsrv.Service {
	gamerRepo := pgrepo.NewGamerRepo(uow)
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	return gamerappsrv.NewService(gamerRepo, domainEventDispatcher)
}

func provideWorldAppService(uow pguow.Uow) worldappsrv.Service {
	worldRepo := pgrepo.NewWorldRepo(uow)
	itemRepo := pgrepo.NewItemRepo(uow)
	unitRepo := pgrepo.NewUnitRepo(uow)
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	return worldappsrv.NewService(worldRepo, unitRepo, itemRepo, domainEventDispatcher)
}
