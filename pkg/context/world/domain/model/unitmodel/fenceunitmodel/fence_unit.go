package fenceunitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

type FenceUnit struct {
	id                   unitmodel.UnitId
	worldId              globalcommonmodel.WorldId
	position             worldcommonmodel.Position
	itemId               worldcommonmodel.ItemId
	direction            worldcommonmodel.Direction
	domainEventCollector *domain.DomainEventCollector
}

// Interface Implementation Check
var _ domain.Aggregate = (*FenceUnit)(nil)

func NewFenceUnit(
	worldId globalcommonmodel.WorldId,
	position worldcommonmodel.Position,
	itemId worldcommonmodel.ItemId,
	direction worldcommonmodel.Direction,
) FenceUnit {
	return FenceUnit{
		id:                   unitmodel.NewUnitId(worldId, position),
		worldId:              worldId,
		position:             position,
		itemId:               itemId,
		direction:            direction,
		domainEventCollector: domain.NewDomainEventCollector(),
	}
}

func LoadFenceUnit(
	id unitmodel.UnitId,
	worldId globalcommonmodel.WorldId,
	position worldcommonmodel.Position,
	itemId worldcommonmodel.ItemId,
	direction worldcommonmodel.Direction,
) FenceUnit {
	return FenceUnit{
		id:                   id,
		worldId:              worldId,
		position:             position,
		itemId:               itemId,
		direction:            direction,
		domainEventCollector: domain.NewDomainEventCollector(),
	}
}

func (unit *FenceUnit) PopDomainEvents() []domain.DomainEvent {
	return unit.domainEventCollector.PopAll()
}

func (unit *FenceUnit) GetId() unitmodel.UnitId {
	return unit.id
}

func (unit *FenceUnit) GetWorldId() globalcommonmodel.WorldId {
	return unit.worldId
}

func (unit *FenceUnit) GetPosition() worldcommonmodel.Position {
	return unit.position
}

func (unit *FenceUnit) GetItemId() worldcommonmodel.ItemId {
	return unit.itemId
}

func (unit *FenceUnit) GetDirection() worldcommonmodel.Direction {
	return unit.direction
}

func (unit *FenceUnit) Rotate() {
	unit.direction = unit.direction.Rotate()
}

func (unit *FenceUnit) Delete() {
}
