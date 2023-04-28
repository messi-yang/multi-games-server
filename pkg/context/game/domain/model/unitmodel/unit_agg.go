package unitmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/domainmodel"
)

type UnitAgg struct {
	worldId      commonmodel.WorldIdVo
	position     commonmodel.PositionVo
	itemId       commonmodel.ItemIdVo
	direction    commonmodel.DirectionVo
	domainEvents []domainmodel.DomainEvent
}

// Interface Implementation Check
var _ domainmodel.Aggregate = (*UnitAgg)(nil)

func NewUnitAgg(
	worldId commonmodel.WorldIdVo,
	position commonmodel.PositionVo,
	itemId commonmodel.ItemIdVo,
	direction commonmodel.DirectionVo,
) UnitAgg {
	return UnitAgg{worldId: worldId, position: position, itemId: itemId, direction: direction, domainEvents: []domainmodel.DomainEvent{}}
}

func (agg *UnitAgg) AddDomainEvent(domainEvent domainmodel.DomainEvent) {
	agg.domainEvents = append(agg.domainEvents, domainEvent)
}

func (agg *UnitAgg) GetDomainEvents() []domainmodel.DomainEvent {
	return agg.domainEvents
}

func (agg *UnitAgg) GetWorldId() commonmodel.WorldIdVo {
	return agg.worldId
}

func (agg *UnitAgg) GetPosition() commonmodel.PositionVo {
	return agg.position
}

func (agg *UnitAgg) GetItemId() commonmodel.ItemIdVo {
	return agg.itemId
}

func (agg *UnitAgg) GetDirection() commonmodel.DirectionVo {
	return agg.direction
}
