package worlddomaineventhandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/application/service/accessappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/model/accessmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/infrastructure/persistence/postgres/pgrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/event/memory/memdomainevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
)

func provideAccessAppService(uow pguow.Uow) accessappsrv.Service {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	worldRoleRepo := pgrepo.NewWorldRoleRepo(uow, domainEventDispatcher)
	accessDomainService := accessmodel.NewAccessService(worldRoleRepo, domainEventDispatcher)
	return accessappsrv.NewService(accessDomainService)
}
