package memdomaineventhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/staticunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/providedependency"
)

type StaticUnitDeletedHandler struct {
}

func NewStaticUnitDeletedHandler() memdomainevent.Handler {
	return &StaticUnitDeletedHandler{}
}

func (handler StaticUnitDeletedHandler) Handle(uow pguow.Uow, domainEvent domain.DomainEvent) error {
	staticUnitDeleted := domainEvent.(staticunitmodel.StaticUnitDeleted)
	unitAppService := providedependency.ProvideUnitAppService(uow)
	return unitAppService.HandleStaticUnitDeletedDomainEvent(staticUnitDeleted)
}
