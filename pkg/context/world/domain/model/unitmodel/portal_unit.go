package unitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/google/uuid"
)

type PortalUnit struct {
	id                   PortalUnitId
	targetPosition       *worldcommonmodel.Position
	domainEventCollector *domain.DomainEventCollector
}

// Interface Implementation Check
var _ domain.Aggregate = (*PortalUnit)(nil)

func NewPortalUnit(
	targetPosition *worldcommonmodel.Position,
) PortalUnit {
	return PortalUnit{
		id:                   NewPortalUnitId(uuid.New()),
		targetPosition:       targetPosition,
		domainEventCollector: domain.NewDomainEventCollector(),
	}
}

func LoadPortalUnit(
	id PortalUnitId,
	targetPosition *worldcommonmodel.Position,
) PortalUnit {
	return PortalUnit{
		id:                   id,
		targetPosition:       targetPosition,
		domainEventCollector: domain.NewDomainEventCollector(),
	}
}

func (portalUnit *PortalUnit) PopDomainEvents() []domain.DomainEvent {
	return portalUnit.domainEventCollector.PopAll()
}

func (portalUnit *PortalUnit) GetId() PortalUnitId {
	return portalUnit.id
}

func (portalUnit *PortalUnit) GetTargetPosition() *worldcommonmodel.Position {
	return portalUnit.targetPosition
}

func (portalUnit *PortalUnit) UpdateTargetPosition(position *worldcommonmodel.Position) {
	portalUnit.targetPosition = position
}
