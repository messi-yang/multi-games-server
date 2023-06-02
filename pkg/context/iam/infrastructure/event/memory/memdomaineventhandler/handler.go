package memdomaineventhandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/infrastructure/event/memory/memdomaineventhandler/worlddomaineventhandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/event/memory/memdomainevent"
)

func Run() {
	domainEventRegister := memdomainevent.NewRegister()
	domainEventRegister.Register(sharedkernelmodel.WorldCreated{}, worlddomaineventhandler.NewWorldCreatedHandler())
}
