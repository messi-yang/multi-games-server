package providedependency

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/worldpermissionappsrv"
)

func ProvideWorldPermissionAppService(uow pguow.Uow) worldpermissionappsrv.Service {
	return worldpermissionappsrv.NewService()
}
