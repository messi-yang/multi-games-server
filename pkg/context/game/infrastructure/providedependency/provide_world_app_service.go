package providedependency

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/application/service/worldappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/service"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/infrastructure/persistence/postgres/pgrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/event/memory/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
)

func ProvideWorldAppService(uow pguow.Uow) worldappsrv.Service {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	gamerRepo := pgrepo.NewGamerRepo(uow, domainEventDispatcher)
	worldRepo := pgrepo.NewWorldRepo(uow, domainEventDispatcher)
	itemRepo := pgrepo.NewItemRepo(uow, domainEventDispatcher)
	unitRepo := pgrepo.NewUnitRepo(uow, domainEventDispatcher)
	worldService := service.NewWorldService(gamerRepo, worldRepo, unitRepo, itemRepo)
	return worldappsrv.NewService(worldRepo, worldService)
}
