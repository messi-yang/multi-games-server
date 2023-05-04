package itemhttphandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/itemappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/unitofwork/pguow"
)

func provideItemAppService(uow pguow.Uow) itemappsrv.Service {
	itemRepo := pgrepo.NewItemRepo(uow)
	return itemappsrv.NewService(itemRepo)
}
