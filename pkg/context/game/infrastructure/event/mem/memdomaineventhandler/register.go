package memdomaineventhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/worldmodel/playermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/worldmodel/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/event/memory/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/messaging/redis/redisservermessagemediator"
)

func RegisterEvents() {
	redisServerMessageMediator := redisservermessagemediator.NewMediator()

	domainEventRegister := memdomainevent.NewRegister()
	domainEventRegister.Register(unitmodel.UnitCreated{}, NewUnitCreatedHandler(redisServerMessageMediator))
	domainEventRegister.Register(unitmodel.UnitDeleted{}, NewUnitDeletedHandler(redisServerMessageMediator))
	domainEventRegister.Register(playermodel.PlayerJoined{}, NewPlayerJoinedHandler(redisServerMessageMediator))
	domainEventRegister.Register(playermodel.PlayerLeft{}, NewPlayerLeftHandler(redisServerMessageMediator))
	domainEventRegister.Register(playermodel.PlayerMoved{}, NewPlayerMovedHandler(redisServerMessageMediator))
	domainEventRegister.Register(sharedkernelmodel.UserCreated{}, NewUserCreatedHandler())
}
