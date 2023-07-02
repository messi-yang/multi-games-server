package unitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/commonmodel"
)

type Unit struct {
	id                   UnitId
	worldId              sharedkernelmodel.WorldId
	position             commonmodel.Position
	itemId               commonmodel.ItemId
	direction            commonmodel.Direction
	domainEventCollector *domain.DomainEventCollector
}

// Interface Implementation Check
var _ domain.Aggregate = (*Unit)(nil)

func NewUnit(
	id UnitId,
	worldId sharedkernelmodel.WorldId,
	position commonmodel.Position,
	itemId commonmodel.ItemId,
	direction commonmodel.Direction,
) Unit {
	unit := Unit{
		id:                   id,
		worldId:              worldId,
		position:             position,
		itemId:               itemId,
		direction:            direction,
		domainEventCollector: domain.NewDomainEventCollector(),
	}
	unit.domainEventCollector.Add(NewUnitCreated(id))
	return unit
}

func LoadUnit(
	id UnitId,
	worldId sharedkernelmodel.WorldId,
	position commonmodel.Position,
	itemId commonmodel.ItemId,
	direction commonmodel.Direction,
) Unit {
	return Unit{
		id:                   id,
		worldId:              worldId,
		position:             position,
		itemId:               itemId,
		direction:            direction,
		domainEventCollector: domain.NewDomainEventCollector(),
	}
}

func (unit *Unit) PopDomainEvents() []domain.DomainEvent {
	return unit.domainEventCollector.PopAll()
}

func (unit *Unit) GetId() UnitId {
	return unit.id
}

func (unit *Unit) GetWorldId() sharedkernelmodel.WorldId {
	return unit.worldId
}

func (unit *Unit) GetPosition() commonmodel.Position {
	return unit.position
}

func (unit *Unit) GetItemId() commonmodel.ItemId {
	return unit.itemId
}

func (unit *Unit) GetDirection() commonmodel.Direction {
	return unit.direction
}

func (unit *Unit) Remove() {
	unit.domainEventCollector.Add(NewUnitDeleted(unit.id))
}
