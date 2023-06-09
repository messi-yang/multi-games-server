package providedependency

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/application/service/accessappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/infrastructure/persistence/postgres/pgrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/event/memory/memdomainevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
)

func ProvideAccessAppService(uow pguow.Uow) accessappsrv.Service {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	userWorldRoleRepo := pgrepo.NewUserWorldRoleRepo(uow, domainEventDispatcher)
	accessDomainService := service.NewAccessService(userWorldRoleRepo, domainEventDispatcher)
	return accessappsrv.NewService(userWorldRoleRepo, accessDomainService)
}
