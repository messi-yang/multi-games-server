package providedependency

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/application/service/worldmemberappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/persistence/pgrepo"
)

func ProvideWorldMemberAppService(uow pguow.Uow) worldmemberappsrv.Service {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	worldMemberRepo := pgrepo.NewWorldMemberRepo(uow, domainEventDispatcher)
	userRepo := pgrepo.NewUserRepo(uow, domainEventDispatcher)
	return worldmemberappsrv.NewService(worldMemberRepo, userRepo)
}
