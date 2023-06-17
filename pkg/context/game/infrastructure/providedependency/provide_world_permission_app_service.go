package providedependency

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/application/service/worldpermissionappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
)

func ProvideWorldPermissionAppService(uow pguow.Uow) worldpermissionappsrv.Service {
	return worldpermissionappsrv.NewService()
}
