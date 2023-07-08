package providedependency

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/application/service/userappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/persistence/pgrepo"
)

func ProvideUserAppService(uow pguow.Uow) userappsrv.Service {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	userRepo := pgrepo.NewUserRepo(uow, domainEventDispatcher)
	return userappsrv.NewService(userRepo)
}
