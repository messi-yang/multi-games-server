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

type PlayerMovedHandler struct {
	redisServerMessageMediator redisservermessagemediator.Mediator
}

func NewPlayerMovedHandler(redisServerMessageMediator redisservermessagemediator.Mediator) memdomainevent.Handler {
	return &PlayerMovedHandler{
		redisServerMessageMediator: redisServerMessageMediator,
	}
}

func (handler PlayerMovedHandler) Handle(uow pguow.Uow, domainEvent domain.DomainEvent) error {
	playerMoved := domainEvent.(playermodel.PlayerMoved)

	uow.AddDelayedWork(func() {
		worldIdDto := playerMoved.GetWorldId().Uuid()
		playerIdDto := playerMoved.GetPlayerId().Uuid()
		handler.redisServerMessageMediator.Send(
			gameappsrv.NewWorldServerMessageChannel(worldIdDto),
			jsonutil.Marshal(gameappsrv.NewPlayerMovedServerMessage(worldIdDto, playerIdDto)),
		)
	})

	return nil
}
