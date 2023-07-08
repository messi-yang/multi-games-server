package providedependency

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/worldpermissionappsrv"
)

func ProvideWorldPermissionAppService(uow pguow.Uow) worldpermissionappsrv.Service {
	return worldpermissionappsrv.NewService()
}
