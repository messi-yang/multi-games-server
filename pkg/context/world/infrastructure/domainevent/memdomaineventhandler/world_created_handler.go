package memdomaineventhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/providedependency"
)

type WorldCreatedHandler struct {
}

func NewWorldCreatedHandler() memdomainevent.Handler {
	return &WorldCreatedHandler{}
}

func (handler WorldCreatedHandler) Handle(uow pguow.Uow, domainEvent domain.DomainEvent) error {
	worldCreated := domainEvent.(worldmodel.WorldCreated)
	worldAccountAppService := providedependency.ProvideWorldAccountAppService(uow)
	return worldAccountAppService.HandleWorldCreatedDomainEvent(worldCreated)
}
