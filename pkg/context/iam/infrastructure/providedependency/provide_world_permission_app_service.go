package providedependency

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/application/service/worldpermissionappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/persistence/pgrepo"
)

func ProvideWorldPermissionAppService(uow pguow.Uow) worldpermissionappsrv.Service {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	worldMemberRepo := pgrepo.NewWorldMemberRepo(uow, domainEventDispatcher)
	return worldpermissionappsrv.NewService(worldMemberRepo)
}
