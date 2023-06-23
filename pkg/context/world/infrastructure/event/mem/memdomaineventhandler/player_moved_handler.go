package memdomaineventhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/event/memory/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/messaging/redis/redisservermessagemediator"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/worldjourneyappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel/playermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/providedependency"
	"github.com/dum-dum-genius/zossi-server/pkg/util/jsonutil"
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
	worldIdDto := playerMoved.GetWorldId().Uuid()
	playerIdDto := playerMoved.GetPlayerId().Uuid()

	worldJourneyAppService := providedependency.ProvideWorldJourneyAppService(uow)
	player, err := worldJourneyAppService.GetPlayer(worldjourneyappsrv.GetPlayerQuery{
		PlayerId: playerIdDto,
	})
	if err != nil {
		return err
	}

	uow.AddDelayedWork(func() {
		handler.redisServerMessageMediator.Send(
			worldjourneyappsrv.NewWorldServerMessageChannel(worldIdDto),
			jsonutil.Marshal(worldjourneyappsrv.NewPlayerMovedServerMessage(player)),
		)
	})

	return nil
}
