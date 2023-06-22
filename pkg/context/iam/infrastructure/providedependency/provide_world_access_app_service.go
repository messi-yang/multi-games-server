package providedependency

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/application/service/worldaccessappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/service"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/persistence/postgres/pgrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/event/memory/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
)

func ProvideWorldAccessAppService(uow pguow.Uow) worldaccessappsrv.Service {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	worldMemberRepo := pgrepo.NewWorldMemberRepo(uow, domainEventDispatcher)
	worldAccessService := service.NewWorldAccessService(worldMemberRepo, domainEventDispatcher)
	return worldaccessappsrv.NewService(worldMemberRepo, worldAccessService)
}
