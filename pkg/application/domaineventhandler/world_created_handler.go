package domaineventhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/domainevent"

	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
)

type WorldCreatedHandler struct{}

func NewWorldCreatedHandler() memdomaineventhandler.Handler {
	return &WorldCreatedHandler{}
}

func ProvideWorldCreatedHandler() memdomaineventhandler.Handler {
	return NewWorldCreatedHandler()
}

func (handler WorldCreatedHandler) Handle(uow pguow.Uow, domainEvent domain.DomainEvent) error {
	worldCreated := domainEvent.(domainevent.WorldCreated)

	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	worldAccountRepo := pgrepo.NewWorldAccountRepo(uow, domainEventDispatcher)

	worldAccount, err := worldAccountRepo.GetWorldAccountOfUser(worldCreated.GetUserId())
	if err != nil {
		return err
	}
	worldAccount.AddWorldsCount()
	return worldAccountRepo.Update(worldAccount)
}
