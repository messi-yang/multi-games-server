package unitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

type Unit struct {
	id                   UnitId
	worldId              globalcommonmodel.WorldId
	position             worldcommonmodel.Position
	itemId               worldcommonmodel.ItemId
	direction            worldcommonmodel.Direction
	domainEventCollector *domain.DomainEventCollector
}

// Interface Implementation Check
var _ domain.Aggregate = (*Unit)(nil)

func NewUnit(
	worldId globalcommonmodel.WorldId,
	position worldcommonmodel.Position,
	itemId worldcommonmodel.ItemId,
	direction worldcommonmodel.Direction,
) Unit {
	return Unit{
		id:                   NewUnitId(worldId, position),
		worldId:              worldId,
		position:             position,
		itemId:               itemId,
		direction:            direction,
		domainEventCollector: domain.NewDomainEventCollector(),
	}
}

func LoadUnit(
	id UnitId,
	worldId globalcommonmodel.WorldId,
	position worldcommonmodel.Position,
	itemId worldcommonmodel.ItemId,
	direction worldcommonmodel.Direction,
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

func (unit *Unit) GetWorldId() globalcommonmodel.WorldId {
	return unit.worldId
}

func (unit *Unit) GetPosition() worldcommonmodel.Position {
	return unit.position
}

func (unit *Unit) GetItemId() worldcommonmodel.ItemId {
	return unit.itemId
}

func (unit *Unit) GetDirection() worldcommonmodel.Direction {
	return unit.direction
}
