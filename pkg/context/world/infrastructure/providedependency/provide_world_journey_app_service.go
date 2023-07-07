package providedependency

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/event/memory/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/worldjourneyappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/service"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/memory/memrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/postgres/pgrepo"
)

func ProvideWorldJourneyAppService(uow pguow.Uow) worldjourneyappsrv.Service {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	itemRepo := pgrepo.NewItemRepo(uow, domainEventDispatcher)
	// playerRepo := pgrepo.NewPlayerRepo(uow, domainEventDispatcher)
	playerRepo := memrepo.NewPlayerRepo(domainEventDispatcher)
	worldRepo := pgrepo.NewWorldRepo(uow, domainEventDispatcher)
	unitRepo := pgrepo.NewUnitRepo(uow, domainEventDispatcher)
	unitService := service.NewUnitService(worldRepo, unitRepo)
	return worldjourneyappsrv.NewService(
		worldRepo,
		playerRepo,
		unitRepo,
		itemRepo,
		unitService,
	)
}
