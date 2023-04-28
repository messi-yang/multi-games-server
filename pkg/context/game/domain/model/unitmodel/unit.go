package unitmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/domainmodel"
)

type Unit struct {
	worldId      commonmodel.WorldId
	position     commonmodel.Position
	itemId       commonmodel.ItemId
	direction    commonmodel.Direction
	domainEvents []domainmodel.DomainEvent
}

// Interface Implementation Check
var _ domainmodel.Aggregate = (*Unit)(nil)

func NewUnit(
	worldId commonmodel.WorldId,
	position commonmodel.Position,
	itemId commonmodel.ItemId,
	direction commonmodel.Direction,
) Unit {
	return Unit{worldId: worldId, position: position, itemId: itemId, direction: direction, domainEvents: []domainmodel.DomainEvent{}}
}

func (unit *Unit) AddDomainEvent(domainEvent domainmodel.DomainEvent) {
	unit.domainEvents = append(unit.domainEvents, domainEvent)
}

func (unit *Unit) GetDomainEvents() []domainmodel.DomainEvent {
	return unit.domainEvents
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
