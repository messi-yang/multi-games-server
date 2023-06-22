package memdomaineventhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/application/service/worldjourneyappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/worldmodel/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/infrastructure/providedependency"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/event/memory/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/messaging/redis/redisservermessagemediator"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/util/jsonutil"
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
	worldIdDto := unitCreated.GetUnitId().GetWorldId().Uuid()
	positionDto := dto.NewPositionDto(unitCreated.GetUnitId().GetPosition())

	worldJourneyAppService := providedependency.ProvideWorldJourneyAppService(uow)
	unitDto, err := worldJourneyAppService.GetUnit(worldjourneyappsrv.GetUnitQuery{
		WorldId:  worldIdDto,
		Position: positionDto,
	})
	if err != nil {
		return err
	}

	uow.AddDelayedWork(func() {
		handler.redisServerMessageMediator.Send(
			worldjourneyappsrv.NewWorldServerMessageChannel(worldIdDto),
			jsonutil.Marshal(worldjourneyappsrv.NewUnitCreatedServerMessage(unitDto)),
		)
	})

	return nil
}
