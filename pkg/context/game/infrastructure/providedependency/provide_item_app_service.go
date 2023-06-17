package providedependency

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/application/service/itemappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/infrastructure/persistence/postgres/pgrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/event/memory/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
)

func ProvideItemAppService(uow pguow.Uow) itemappsrv.Service {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	itemRepo := pgrepo.NewItemRepo(uow, domainEventDispatcher)
	return itemappsrv.NewService(itemRepo)
}
