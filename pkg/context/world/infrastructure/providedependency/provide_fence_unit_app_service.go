package providedependency

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/fenceunitappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/service"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
)

func ProvideFenceUnitAppService(uow pguow.Uow) fenceunitappsrv.Service {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	itemRepo := pgrepo.NewItemRepo(uow, domainEventDispatcher)
	worldRepo := pgrepo.NewWorldRepo(uow, domainEventDispatcher)
	unitRepo := pgrepo.NewUnitRepo(uow, domainEventDispatcher)
	fenceUnitRepo := pgrepo.NewFenceUnitRepo(uow, domainEventDispatcher)
	fenceUnitRepoUnitService := service.NewFenceUnitService(worldRepo, unitRepo, fenceUnitRepo, itemRepo)
	return fenceunitappsrv.NewService(
		fenceUnitRepo,
		fenceUnitRepoUnitService,
	)
}
