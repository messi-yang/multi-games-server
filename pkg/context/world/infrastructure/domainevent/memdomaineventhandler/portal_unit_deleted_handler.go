package memdomaineventhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/portalunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/providedependency"
)

type PortalUnitDeletedHandler struct {
}

func NewPortalUnitDeletedHandler() memdomainevent.Handler {
	return &PortalUnitDeletedHandler{}
}

func (handler PortalUnitDeletedHandler) Handle(uow pguow.Uow, domainEvent domain.DomainEvent) error {
	portalUnitDeleted := domainEvent.(portalunitmodel.PortalUnitDeleted)
	unitAppService := providedependency.ProvideUnitAppService(uow)
	return unitAppService.HandlePortalUnitDeletedDomainEvent(portalUnitDeleted)
}
