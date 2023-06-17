package providedependency

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/application/service/userappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/persistence/postgres/pgrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/event/memory/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
)

func ProvideUserAppService(uow pguow.Uow) userappsrv.Service {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	userRepo := pgrepo.NewUserRepo(uow, domainEventDispatcher)
	return userappsrv.NewService(userRepo)
}
