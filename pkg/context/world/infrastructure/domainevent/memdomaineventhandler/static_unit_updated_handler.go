package memdomaineventhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/staticunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/providedependency"
)

type StaticUnitUpdatedHandler struct {
}

func NewStaticUnitUpdatedHandler() memdomainevent.Handler {
	return &StaticUnitUpdatedHandler{}
}

func (handler StaticUnitUpdatedHandler) Handle(uow pguow.Uow, domainEvent domain.DomainEvent) error {
	staticUnitUpdated := domainEvent.(staticunitmodel.StaticUnitUpdated)
	unitAppService := providedependency.ProvideUnitAppService(uow)
	return unitAppService.HandleStaticUnitUpdatedDomainEvent(staticUnitUpdated)
}
