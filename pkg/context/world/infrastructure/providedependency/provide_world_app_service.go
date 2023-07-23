package providedependency

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	global_pg_repo "github.com/dum-dum-genius/zossi-server/pkg/context/global/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/worldappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/service"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
)

func ProvideWorldAppService(uow pguow.Uow) worldappsrv.Service {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	worldAccountRepo := pgrepo.NewWorldAccountRepo(uow, domainEventDispatcher)
	worldRepo := pgrepo.NewWorldRepo(uow, domainEventDispatcher)
	itemRepo := pgrepo.NewItemRepo(uow, domainEventDispatcher)
	unitRepo := pgrepo.NewUnitRepo(uow, domainEventDispatcher)
	userRepo := global_pg_repo.NewUserRepo(uow, domainEventDispatcher)
	worldService := service.NewWorldService(worldAccountRepo, worldRepo, unitRepo, itemRepo)
	return worldappsrv.NewService(worldRepo, userRepo, worldService)
}
