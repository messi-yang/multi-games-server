package unitdomaineventhandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/dto"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/gameappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/event/memory/memdomainevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/messaging/redis/redisservermessagemediator"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/util/jsonutil"
)

type UnitCreatedHandler struct {
	redisServerMessageMediator redisservermessagemediator.Mediator
}

func NewUnitCreatedHandler(redisServerMessageMediator redisservermessagemediator.Mediator) memdomainevent.Handler {
	return &UnitCreatedHandler{
		redisServerMessageMediator: redisServerMessageMediator,
	}
}

func (handler UnitCreatedHandler) Handle(uow pguow.Uow, domainEvent domain.DomainEvent) error {
	unitCreated := domainEvent.(unitmodel.UnitCreated)

	uow.AddDelayedWork(func() {
		worldIdDto := unitCreated.GetUnitId().GetWorldId().Uuid()
		positionDto := dto.NewPositionDto(unitCreated.GetUnitId().GetPosition())
		handler.redisServerMessageMediator.Send(
			gameappsrv.NewWorldServerMessageChannel(worldIdDto),
			jsonutil.Marshal(gameappsrv.NewUnitCreatedServerMessage(worldIdDto, positionDto)),
		)
	})

	return nil
}
