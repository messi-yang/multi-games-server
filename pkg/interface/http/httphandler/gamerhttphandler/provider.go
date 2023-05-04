package gamerhttphandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/gamerappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/unitofwork/pguow"
)

func provideGamerAppService(uow pguow.Uow) gamerappsrv.Service {
	gamerRepo := pgrepo.NewGamerRepo(uow)
	return gamerappsrv.NewService(gamerRepo)
}
