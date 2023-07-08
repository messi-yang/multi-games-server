package providedependency

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/application/service/worldaccessappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/service"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/persistence/pgrepo"
)

func ProvideWorldAccessAppService(uow pguow.Uow) worldaccessappsrv.Service {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	worldMemberRepo := pgrepo.NewWorldMemberRepo(uow, domainEventDispatcher)
	worldAccessService := service.NewWorldAccessService(worldMemberRepo, domainEventDispatcher)
	return worldaccessappsrv.NewService(worldMemberRepo, worldAccessService)
}
