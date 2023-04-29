package itemhttphandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/itemappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/pguow"
)

func provideItemAppService(pgUow *pguow.Uow) itemappsrv.Service {
	itemRepo := pgrepo.NewItemRepo(pgUow)
	return itemappsrv.NewService(itemRepo)
}
