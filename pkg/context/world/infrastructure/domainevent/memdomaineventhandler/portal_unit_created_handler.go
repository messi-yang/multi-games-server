package memdomaineventhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/providedependency"
)

type PortalUnitCreatedHandler struct {
}

func NewPortalUnitCreatedHandler() memdomainevent.Handler {
	return &PortalUnitCreatedHandler{}
}

func (handler PortalUnitCreatedHandler) Handle(uow pguow.Uow, domainEvent domain.DomainEvent) error {
	portalUnitCreated := domainEvent.(unitmodel.PortalUnitCreated)
	unitAppService := providedependency.ProvideUnitAppService(uow)
	return unitAppService.HandlePortalUnitCreatedDomainEvent(portalUnitCreated)
}
