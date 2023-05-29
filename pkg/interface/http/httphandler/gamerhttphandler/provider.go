package gamerhttphandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/gamerappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/postgres/pgrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/event/memory/memdomainevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
)

func provideGamerAppService(uow pguow.Uow) gamerappsrv.Service {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	gamerRepo := pgrepo.NewGamerRepo(uow, domainEventDispatcher)
	return gamerappsrv.NewService(gamerRepo)
}
