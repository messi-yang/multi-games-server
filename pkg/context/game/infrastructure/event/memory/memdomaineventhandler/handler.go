package memdomaineventhandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/worldmodel/playermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/worldmodel/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/event/memory/memdomaineventhandler/playerdomaineventhandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/event/memory/memdomaineventhandler/unitdomaineventhandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/event/memory/memdomaineventhandler/userdomaineventhandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/event/memory/memdomainevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/messaging/redis/redisservermessagemediator"
)

func Run() {
	redisServerMessageMediator := redisservermessagemediator.NewMediator()

	domainEventRegister := memdomainevent.NewRegister()
	domainEventRegister.Register(unitmodel.UnitCreated{}, unitdomaineventhandler.NewUnitCreatedHandler(redisServerMessageMediator))
	domainEventRegister.Register(unitmodel.UnitDeleted{}, unitdomaineventhandler.NewUnitDeletedHandler(redisServerMessageMediator))
	domainEventRegister.Register(playermodel.PlayerJoined{}, playerdomaineventhandler.NewPlayerJoinedHandler(redisServerMessageMediator))
	domainEventRegister.Register(playermodel.PlayerLeft{}, playerdomaineventhandler.NewPlayerLeftHandler(redisServerMessageMediator))
	domainEventRegister.Register(playermodel.PlayerMoved{}, playerdomaineventhandler.NewPlayerMovedHandler(redisServerMessageMediator))
	domainEventRegister.Register(sharedkernelmodel.UserCreated{}, userdomaineventhandler.NewUserCreatedHandler())
}
