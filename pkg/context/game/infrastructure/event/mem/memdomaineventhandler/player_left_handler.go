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

type PlayerLeftHandler struct {
	redisServerMessageMediator redisservermessagemediator.Mediator
}

func NewPlayerLeftHandler(redisServerMessageMediator redisservermessagemediator.Mediator) memdomainevent.Handler {
	return &PlayerLeftHandler{
		redisServerMessageMediator: redisServerMessageMediator,
	}
}

func (handler PlayerLeftHandler) Handle(uow pguow.Uow, domainEvent domain.DomainEvent) error {
	playerLeft := domainEvent.(playermodel.PlayerLeft)

	uow.AddDelayedWork(func() {
		worldIdDto := playerLeft.GetWorldId().Uuid()
		playerIdDto := playerLeft.GetPlayerId().Uuid()
		handler.redisServerMessageMediator.Send(
			gameappsrv.NewWorldServerMessageChannel(worldIdDto),
			jsonutil.Marshal(gameappsrv.NewPlayerLeftServerMessage(worldIdDto, playerIdDto)),
		)
	})

	return nil
}
