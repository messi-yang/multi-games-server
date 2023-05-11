package unitmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
)

type Unit struct {
	id                   commonmodel.UnitId
	worldId              commonmodel.WorldId
	position             commonmodel.Position
	itemId               commonmodel.ItemId
	direction            commonmodel.Direction
	domainEventCollector *domain.DomainEventCollector
}

// Interface Implementation Check
var _ domain.Aggregate = (*Unit)(nil)

func NewUnit(
	id commonmodel.UnitId,
	worldId commonmodel.WorldId,
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
	id commonmodel.UnitId,
	worldId commonmodel.WorldId,
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

func (unit *Unit) GetId() commonmodel.UnitId {
	return unit.id
}

func (unit *Unit) GetWorldId() commonmodel.WorldId {
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

func (unit *Unit) Delete() {
	unit.domainEventCollector.Add(NewUnitDeleted(unit.id))
}
