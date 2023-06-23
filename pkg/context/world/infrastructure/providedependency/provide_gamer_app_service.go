package providedependency

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/event/memory/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/worldaccountappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/postgres/pgrepo"
)

func ProvideWorldAccountAppService(uow pguow.Uow) worldaccountappsrv.Service {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	worldAccountRepo := pgrepo.NewWorldAccountRepo(uow, domainEventDispatcher)
	return worldaccountappsrv.NewService(worldAccountRepo)
}
