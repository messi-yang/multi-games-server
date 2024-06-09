package domaineventhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/domainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
)

type WorldDeletedHandler struct{}

func NewWorldDeletedHandler() memdomaineventhandler.Handler {
	return &WorldDeletedHandler{}
}

func ProvideWorldDeletedHandler() memdomaineventhandler.Handler {
	return NewWorldDeletedHandler()
}

func (handler WorldDeletedHandler) Handle(uow pguow.Uow, domainEvent domain.DomainEvent) error {
	worldDeleted := domainEvent.(domainevent.WorldDeleted)

	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	worldAccountRepo := pgrepo.NewWorldAccountRepo(uow, domainEventDispatcher)

	worldAccount, err := worldAccountRepo.GetWorldAccountOfUser(worldDeleted.GetUserId())
	if err != nil {
		return err
	}
	worldAccount.SubtractWorldsCount()
	return worldAccountRepo.Update(worldAccount)
}
