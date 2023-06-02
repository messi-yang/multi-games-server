package memdomaineventhandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/gameappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/worldmodel/playermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/event/memory/memdomainevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/messaging/redis/redisservermessagemediator"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/util/jsonutil"
)

type PlayerJoinedHandler struct {
	redisServerMessageMediator redisservermessagemediator.Mediator
}

func NewPlayerJoinedHandler(redisServerMessageMediator redisservermessagemediator.Mediator) memdomainevent.Handler {
	return &PlayerJoinedHandler{
		redisServerMessageMediator: redisServerMessageMediator,
	}
}

func (handler PlayerJoinedHandler) Handle(uow pguow.Uow, domainEvent domain.DomainEvent) error {
	playerJoined := domainEvent.(playermodel.PlayerJoined)

	uow.AddDelayedWork(func() {
		worldIdDto := playerJoined.GetWorldId().Uuid()
		playerIdDto := playerJoined.GetPlayerId().Uuid()
		handler.redisServerMessageMediator.Send(
			gameappsrv.NewWorldServerMessageChannel(worldIdDto),
			jsonutil.Marshal(gameappsrv.NewPlayerJoinedServerMessage(worldIdDto, playerIdDto)),
		)
	})

	return nil
}
