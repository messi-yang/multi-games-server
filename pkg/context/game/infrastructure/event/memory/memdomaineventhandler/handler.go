package memdomaineventhandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/event/memory/memdomaineventhandler/unitdomaineventhandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/event/memory/memdomainevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/messaging/redis/redisservermessagemediator"
)

func Run() {
	redisServerMessageMediator := redisservermessagemediator.NewMediator()

	domainEventRegister := memdomainevent.NewRegister()
	domainEventRegister.Register(unitmodel.UnitCreated{}, unitdomaineventhandler.NewUnitCreatedHandler(redisServerMessageMediator))
	domainEventRegister.Register(unitmodel.UnitDeleted{}, unitdomaineventhandler.NewUnitDeletedHandler(redisServerMessageMediator))
}
