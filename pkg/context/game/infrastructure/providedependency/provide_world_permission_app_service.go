package providedependency

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/worldpermissionappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
)

func ProvideWorldPermissionAppService(uow pguow.Uow) worldpermissionappsrv.Service {
	return worldpermissionappsrv.NewService()
}
