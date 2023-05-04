package unitmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
)

type Unit struct {
	worldId              commonmodel.WorldId
	position             commonmodel.Position
	itemId               commonmodel.ItemId
	direction            commonmodel.Direction
	domainEventCollector *domain.DomainEventCollector
}

// Interface Implementation Check
var _ domain.Aggregate = (*Unit)(nil)

func NewUnit(
	worldId commonmodel.WorldId,
	position commonmodel.Position,
	itemId commonmodel.ItemId,
	direction commonmodel.Direction,
) Unit {
	return Unit{
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
