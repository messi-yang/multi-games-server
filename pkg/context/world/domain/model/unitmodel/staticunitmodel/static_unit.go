package staticunitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

type StaticUnit struct {
	id                   StaticUnitId
	worldId              globalcommonmodel.WorldId
	position             worldcommonmodel.Position
	itemId               worldcommonmodel.ItemId
	direction            worldcommonmodel.Direction
	domainEventCollector *domain.DomainEventCollector
}

// Interface Implementation Check
var _ domain.Aggregate = (*StaticUnit)(nil)

func NewStaticUnit(
	id StaticUnitId,
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

func LoadStaticUnit(
	id StaticUnitId,
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

func (staticUnit *StaticUnit) PopDomainEvents() []domain.DomainEvent {
	return staticUnit.domainEventCollector.PopAll()
}

func (staticUnit *StaticUnit) GetId() StaticUnitId {
	return staticUnit.id
}

func (staticUnit *StaticUnit) GetWorldId() globalcommonmodel.WorldId {
	return staticUnit.worldId
}

func (staticUnit *StaticUnit) GetPosition() worldcommonmodel.Position {
	return staticUnit.position
}

func (staticUnit *StaticUnit) GetItemId() worldcommonmodel.ItemId {
	return staticUnit.itemId
}

func (staticUnit *StaticUnit) GetDirection() worldcommonmodel.Direction {
	return staticUnit.direction
}

func (staticUnit *StaticUnit) Rotate() {
	staticUnit.direction = staticUnit.direction.Rotate()
}

func (staticUnit *StaticUnit) Delete() {
}
