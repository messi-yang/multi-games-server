package providedependency

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/worldaccountappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
)

func ProvideWorldAccountAppService(uow pguow.Uow) worldaccountappsrv.Service {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	worldAccountRepo := pgrepo.NewWorldAccountRepo(uow, domainEventDispatcher)
	return worldaccountappsrv.NewService(worldAccountRepo)
}
