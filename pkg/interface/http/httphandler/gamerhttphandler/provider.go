package gamerhttphandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/gamerappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/pguow"
)

func provideGamerAppService(pgUow pguow.Uow) gamerappsrv.Service {
	gamerRepo := pgrepo.NewGamerRepo(pgUow)
	return gamerappsrv.NewService(gamerRepo)
}
