package memdomaineventhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/event/memory/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/worldaccountappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/providedependency"
)

type UserCreatedHandler struct {
}

func NewUserCreatedHandler() memdomainevent.Handler {
	return &UserCreatedHandler{}
}

func (handler UserCreatedHandler) Handle(uow pguow.Uow, domainEvent domain.DomainEvent) error {
	userCreated := domainEvent.(sharedkernelmodel.UserCreated)
	worldAccountAppService := providedependency.ProvideWorldAccountAppService(uow)

	if _, err := worldAccountAppService.CreateWorldAccount(worldaccountappsrv.CreateWorldAccountCommand{
		UserId: userCreated.GetUserId().Uuid(),
	}); err != nil {
		return err
	}
	return nil
}
