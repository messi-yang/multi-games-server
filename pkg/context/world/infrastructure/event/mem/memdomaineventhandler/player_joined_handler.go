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
	worldIdDto := playerJoined.GetWorldId().Uuid()
	playerIdDto := playerJoined.GetPlayerId().Uuid()

	worldJourneyAppService := providedependency.ProvideWorldJourneyAppService(uow)
	player, err := worldJourneyAppService.GetPlayer(worldjourneyappsrv.GetPlayerQuery{
		WorldId:  worldIdDto,
		PlayerId: playerIdDto,
	})
	if err != nil {
		return err
	}

	uow.AddDelayedWork(func() {
		handler.redisServerMessageMediator.Send(
			worldjourneyappsrv.NewWorldServerMessageChannel(worldIdDto),
			jsonutil.Marshal(worldjourneyappsrv.NewPlayerJoinedServerMessage(player)),
		)
	})

	return nil
}
