package providedependency

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/playerappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/redisrepo"
)

func ProvidePlayerAppService() playerappsrv.Service {
	uow := pguow.NewDummyUow()
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	playerRepo := redisrepo.NewPlayerRepo(domainEventDispatcher)
	unitRepo := pgrepo.NewUnitRepo(uow, domainEventDispatcher)
	portalUnitRepo := pgrepo.NewPortalUnitRepo(uow, domainEventDispatcher)
	return playerappsrv.NewService(
		playerRepo,
		unitRepo,
		portalUnitRepo,
	)
}
