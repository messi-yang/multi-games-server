package providedependency

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/worldappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
)

func ProvideWorldAppService(uow pguow.Uow) worldappsrv.Service {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	worldRepo := pgrepo.NewWorldRepo(uow, domainEventDispatcher)
	return worldappsrv.NewService(worldRepo)
}
