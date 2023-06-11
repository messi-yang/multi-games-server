package providedependency

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/application/service/worldaccessappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/infrastructure/persistence/postgres/pgrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/event/memory/memdomainevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
)

func ProvideWorldAccessAppService(uow pguow.Uow) worldaccessappsrv.Service {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	userWorldRoleRepo := pgrepo.NewUserWorldRoleRepo(uow, domainEventDispatcher)
	worldAccessDomainService := service.NewWorldAccessService(userWorldRoleRepo, domainEventDispatcher)
	return worldaccessappsrv.NewService(userWorldRoleRepo, worldAccessDomainService)
}
