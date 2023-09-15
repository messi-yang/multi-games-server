package portalunitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/google/uuid"
)

type PortalUnit struct {
	id                   PortalUnitId
	worldId              globalcommonmodel.WorldId
	position             worldcommonmodel.Position
	itemId               worldcommonmodel.ItemId
	direction            worldcommonmodel.Direction
	targetPosition       *worldcommonmodel.Position
	domainEventCollector *domain.DomainEventCollector
}

// Interface Implementation Check
var _ domain.Aggregate = (*PortalUnit)(nil)

func NewPortalUnit(
	worldId globalcommonmodel.WorldId,
	position worldcommonmodel.Position,
	itemId worldcommonmodel.ItemId,
	direction worldcommonmodel.Direction,
	targetPosition *worldcommonmodel.Position,
) PortalUnit {
	portalUnit := PortalUnit{
		id:                   NewPortalUnitId(uuid.New()),
		worldId:              worldId,
		position:             position,
		itemId:               itemId,
		direction:            direction,
		targetPosition:       targetPosition,
		domainEventCollector: domain.NewDomainEventCollector(),
	}
	portalUnit.domainEventCollector.Add(NewPortalUnitCreated(portalUnit))
	return portalUnit
}

func LoadPortalUnit(
	id PortalUnitId,
	worldId globalcommonmodel.WorldId,
	position worldcommonmodel.Position,
	itemId worldcommonmodel.ItemId,
	direction worldcommonmodel.Direction,
	targetPosition *worldcommonmodel.Position,
) PortalUnit {
	return PortalUnit{
		id:                   id,
		worldId:              worldId,
		position:             position,
		itemId:               itemId,
		direction:            direction,
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

func (portalUnit *PortalUnit) GetWorldId() globalcommonmodel.WorldId {
	return portalUnit.worldId
}

func (portalUnit *PortalUnit) GetPosition() worldcommonmodel.Position {
	return portalUnit.position
}

func (portalUnit *PortalUnit) GetItemId() worldcommonmodel.ItemId {
	return portalUnit.itemId
}

func (portalUnit *PortalUnit) GetDirection() worldcommonmodel.Direction {
	return portalUnit.direction
}

func (portalUnit *PortalUnit) GetTargetPosition() *worldcommonmodel.Position {
	return portalUnit.targetPosition
}

func (portalUnit *PortalUnit) UpdateTargetPosition(targetPosition *worldcommonmodel.Position) {
	portalUnit.targetPosition = targetPosition
}

func (portalUnit *PortalUnit) Rotate() {
	portalUnit.direction = portalUnit.direction.Rotate()
	portalUnit.domainEventCollector.Add(NewPortalUnitUpdated(*portalUnit))
}

func (portalUnit *PortalUnit) Delete() {
	portalUnit.domainEventCollector.Add(NewPortalUnitDeleted(*portalUnit))
}
