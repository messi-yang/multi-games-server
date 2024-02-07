package providedependency

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/linkunitappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/service"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
)

func ProvideLinkUnitAppService(uow pguow.Uow) linkunitappsrv.Service {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	itemRepo := pgrepo.NewItemRepo(uow, domainEventDispatcher)
	worldRepo := pgrepo.NewWorldRepo(uow, domainEventDispatcher)
	unitRepo := pgrepo.NewUnitRepo(uow, domainEventDispatcher)
	linkUnitRepo := pgrepo.NewLinkUnitRepo(uow, domainEventDispatcher)
	linkUnitService := service.NewLinkUnitService(worldRepo, unitRepo, linkUnitRepo, itemRepo)
	return linkunitappsrv.NewService(
		linkUnitRepo,
		linkUnitService,
	)
}
