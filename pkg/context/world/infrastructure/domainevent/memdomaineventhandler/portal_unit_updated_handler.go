package memdomaineventhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/portalunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/providedependency"
)

type PortalUnitUpdatedHandler struct {
}

func NewPortalUnitUpdatedHandler() memdomainevent.Handler {
	return &PortalUnitUpdatedHandler{}
}

func (handler PortalUnitUpdatedHandler) Handle(uow pguow.Uow, domainEvent domain.DomainEvent) error {
	portalUnitUpdated := domainEvent.(portalunitmodel.PortalUnitUpdated)
	unitAppService := providedependency.ProvideUnitAppService(uow)
	return unitAppService.HandlePortalUnitUpdatedDomainEvent(portalUnitUpdated)
}
