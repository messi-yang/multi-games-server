package memdomaineventhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/staticunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/providedependency"
)

type StaticUnitCreatedHandler struct {
}

func NewStaticUnitCreatedHandler() memdomainevent.Handler {
	return &StaticUnitCreatedHandler{}
}

func (handler StaticUnitCreatedHandler) Handle(uow pguow.Uow, domainEvent domain.DomainEvent) error {
	staticUnitCreated := domainEvent.(staticunitmodel.StaticUnitCreated)
	unitAppService := providedependency.ProvideUnitAppService(uow)
	return unitAppService.HandleStaticUnitCreatedDomainEvent(staticUnitCreated)
}
