package worldhttphandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/worldappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/postgres/pgrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/event/memory/memdomainevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
)

func provideWorldAppService(uow pguow.Uow) worldappsrv.Service {
	worldRepo := pgrepo.NewWorldRepo(uow)
	itemRepo := pgrepo.NewItemRepo(uow)
	unitRepo := pgrepo.NewUnitRepo(uow)
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	return worldappsrv.NewService(worldRepo, unitRepo, itemRepo, domainEventDispatcher)
}
