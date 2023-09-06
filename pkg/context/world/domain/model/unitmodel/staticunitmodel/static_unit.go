package staticunitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

type StaticUnit struct {
	id                   unitmodel.UnitId
	worldId              globalcommonmodel.WorldId
	position             worldcommonmodel.Position
	itemId               worldcommonmodel.ItemId
	direction            worldcommonmodel.Direction
	domainEventCollector *domain.DomainEventCollector
}

// Interface Implementation Check
var _ domain.Aggregate = (*StaticUnit)(nil)

func NewStaticUnit(
	worldId globalcommonmodel.WorldId,
	position worldcommonmodel.Position,
	itemId worldcommonmodel.ItemId,
	direction worldcommonmodel.Direction,
) StaticUnit {
	portalUnit := StaticUnit{
		id:                   unitmodel.NewUnitId(worldId, position),
		worldId:              worldId,
		position:             position,
		itemId:               itemId,
		direction:            direction,
		domainEventCollector: domain.NewDomainEventCollector(),
	}
	portalUnit.domainEventCollector.Add(NewStaticUnitCreated(portalUnit))
	return portalUnit
}

func LoadStaticUnit(
	id unitmodel.UnitId,
	worldId globalcommonmodel.WorldId,
	position worldcommonmodel.Position,
	itemId worldcommonmodel.ItemId,
	direction worldcommonmodel.Direction,
) StaticUnit {
	return StaticUnit{
		id:                   id,
		worldId:              worldId,
		position:             position,
		itemId:               itemId,
		direction:            direction,
		domainEventCollector: domain.NewDomainEventCollector(),
	}
}

func (portalUnit *StaticUnit) PopDomainEvents() []domain.DomainEvent {
	return portalUnit.domainEventCollector.PopAll()
}

func (portalUnit *StaticUnit) GetId() unitmodel.UnitId {
	return portalUnit.id
}

func (portalUnit *StaticUnit) GetWorldId() globalcommonmodel.WorldId {
	return portalUnit.worldId
}

func (portalUnit *StaticUnit) GetPosition() worldcommonmodel.Position {
	return portalUnit.position
}

func (portalUnit *StaticUnit) GetItemId() worldcommonmodel.ItemId {
	return portalUnit.itemId
}

func (portalUnit *StaticUnit) GetDirection() worldcommonmodel.Direction {
	return portalUnit.direction
}

func (portalUnit *StaticUnit) Delete() {
	portalUnit.domainEventCollector.Add(NewStaticUnitDeleted(*portalUnit))
}
