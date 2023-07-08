package providedependency

import (
	"os"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/application/service/authappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/service"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/persistence/pgrepo"
)

func ProvideAuthAppService(uow pguow.Uow) authappsrv.Service {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	userRepo := pgrepo.NewUserRepo(uow, domainEventDispatcher)
	identityDomainService := service.NewIdentityService(userRepo)
	return authappsrv.NewService(userRepo, identityDomainService, os.Getenv("AUTH_SECRET"))
}
