package memdomaineventhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/application/service/gameappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/worldmodel/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/event/memory/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/messaging/redis/redisservermessagemediator"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/util/jsonutil"
)

type UnitDeletedHandler struct {
	redisServerMessageMediator redisservermessagemediator.Mediator
}

func NewUnitDeletedHandler(redisServerMessageMediator redisservermessagemediator.Mediator) memdomainevent.Handler {
	return &UnitDeletedHandler{
		redisServerMessageMediator: redisServerMessageMediator,
	}
}

func (handler UnitDeletedHandler) Handle(uow pguow.Uow, domainEvent domain.DomainEvent) error {
	unitDeleted := domainEvent.(unitmodel.UnitDeleted)
	worldIdDto := unitDeleted.GetUnitId().GetWorldId().Uuid()
	positionDto := dto.NewPositionDto(unitDeleted.GetUnitId().GetPosition())

	uow.AddDelayedWork(func() {
		handler.redisServerMessageMediator.Send(
			gameappsrv.NewWorldServerMessageChannel(worldIdDto),
			jsonutil.Marshal(gameappsrv.NewUnitDeletedServerMessage(worldIdDto, positionDto)),
		)
	})

	return nil
}
