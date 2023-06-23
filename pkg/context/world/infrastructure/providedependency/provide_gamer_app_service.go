package providedependency

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/event/memory/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/gamerappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/postgres/pgrepo"
)

func ProvideGamerAppService(uow pguow.Uow) gamerappsrv.Service {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	gamerRepo := pgrepo.NewGamerRepo(uow, domainEventDispatcher)
	return gamerappsrv.NewService(gamerRepo)
}
