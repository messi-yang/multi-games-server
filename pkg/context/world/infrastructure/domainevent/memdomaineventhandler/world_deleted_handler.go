package memdomaineventhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/providedependency"
)

type WorldDeletedHandler struct {
}

func NewWorldDeletedHandler() memdomainevent.Handler {
	return &WorldDeletedHandler{}
}

func (handler WorldDeletedHandler) Handle(uow pguow.Uow, domainEvent domain.DomainEvent) error {
	worldDeleted := domainEvent.(worldmodel.WorldDeleted)
	worldAccountAppService := providedependency.ProvideWorldAccountAppService(uow)
	return worldAccountAppService.HandleWorldDeletedDomainEvent(worldDeleted)
}
