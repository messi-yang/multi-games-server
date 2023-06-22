package memdomaineventhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/application/service/gameappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/worldmodel/playermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/infrastructure/providedependency"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/event/memory/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/messaging/redis/redisservermessagemediator"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
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

	gameAppService := providedependency.ProvideGameAppService(uow)
	player, err := gameAppService.GetPlayer(gameappsrv.GetPlayerQuery{
		PlayerId: playerIdDto,
	})
	if err != nil {
		return err
	}

	uow.AddDelayedWork(func() {
		handler.redisServerMessageMediator.Send(
			gameappsrv.NewWorldServerMessageChannel(worldIdDto),
			jsonutil.Marshal(gameappsrv.NewPlayerMovedServerMessage(player)),
		)
	})

	return nil
}
