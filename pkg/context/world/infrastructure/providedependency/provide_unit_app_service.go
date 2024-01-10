package providedependency

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/unitappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/service"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
)

func ProvideUnitAppService(uow pguow.Uow) unitappsrv.Service {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	itemRepo := pgrepo.NewItemRepo(uow, domainEventDispatcher)
	worldRepo := pgrepo.NewWorldRepo(uow, domainEventDispatcher)
	unitRepo := pgrepo.NewUnitRepo(uow, domainEventDispatcher)
	portalUnitRepo := pgrepo.NewPortalUnitRepo(uow, domainEventDispatcher)
	staticUnitRepo := pgrepo.NewStaticUnitRepo(uow, domainEventDispatcher)
	fenceUnitRepo := pgrepo.NewFenceUnitRepo(uow, domainEventDispatcher)
	staticUnitService := service.NewStaticUnitService(worldRepo, unitRepo, staticUnitRepo, itemRepo)
	fenceUnitService := service.NewFenceUnitService(worldRepo, unitRepo, fenceUnitRepo, itemRepo)
	portalUnitService := service.NewPortalUnitService(worldRepo, unitRepo, portalUnitRepo, itemRepo)
	return unitappsrv.NewService(
		worldRepo,
		unitRepo,
		itemRepo,
		staticUnitService,
		fenceUnitService,
		portalUnitService,
	)
}
