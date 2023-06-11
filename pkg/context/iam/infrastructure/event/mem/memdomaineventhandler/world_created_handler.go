package memdomaineventhandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/application/service/worldaccessappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/infrastructure/providedependency"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/event/memory/memdomainevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
)

type WorldCreatedHandler struct {
}

func NewWorldCreatedHandler() memdomainevent.Handler {
	return &WorldCreatedHandler{}
}

func (handler WorldCreatedHandler) Handle(uow pguow.Uow, domainEvent domain.DomainEvent) error {
	worldCreated := domainEvent.(sharedkernelmodel.WorldCreated)

	worldAccessAppService := providedependency.ProvideWorldAccessAppService(uow)

	return worldAccessAppService.AssignWorldRoleToUser(worldaccessappsrv.AssignWorldRoleToUserCommand{
		UserId:    worldCreated.GetUserId().Uuid(),
		WorldId:   worldCreated.GetWorldId().Uuid(),
		WorldRole: "owner",
	})
}
