package providedependency

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/application/service/worldjourneyappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/service"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/infrastructure/persistence/postgres/pgrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/event/memory/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
)

func ProvideWorldJourneyAppService(uow pguow.Uow) worldjourneyappsrv.Service {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	itemRepo := pgrepo.NewItemRepo(uow, domainEventDispatcher)
	playerRepo := pgrepo.NewPlayerRepo(uow, domainEventDispatcher)
	worldRepo := pgrepo.NewWorldRepo(uow, domainEventDispatcher)
	unitRepo := pgrepo.NewUnitRepo(uow, domainEventDispatcher)
	worldJourneyService := service.NewWorldJourneyService(worldRepo, playerRepo, unitRepo, itemRepo)
	return worldjourneyappsrv.NewService(
		worldRepo, playerRepo, unitRepo, itemRepo, worldJourneyService,
	)
}
