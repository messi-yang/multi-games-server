package unitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/google/uuid"
)

// Unit here is only for reading purpose, for writing units,
// please check the unit model of the type you are looking for.
type Unit struct {
	id                   UnitId
	worldId              globalcommonmodel.WorldId
	position             worldcommonmodel.Position
	itemId               worldcommonmodel.ItemId
	direction            worldcommonmodel.Direction
	_type                worldcommonmodel.UnitType
	info                 any
	infoId               uuid.UUID
	domainEventCollector *domain.DomainEventCollector
}

// Interface Implementation Check
var _ domain.Aggregate = (*Unit)(nil)

func LoadUnit(
	id UnitId,
	worldId globalcommonmodel.WorldId,
	position worldcommonmodel.Position,
	itemId worldcommonmodel.ItemId,
	direction worldcommonmodel.Direction,
	_type worldcommonmodel.UnitType,
	info any,
	infoId uuid.UUID,
) Unit {
	return Unit{
		id:                   id,
		worldId:              worldId,
		position:             position,
		itemId:               itemId,
		direction:            direction,
		_type:                _type,
		info:                 info,
		infoId:               infoId,
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

func (unit *Unit) GetInfo() any {
	return unit.info
}

func (unit *Unit) GetInfoId() uuid.UUID {
	return unit.infoId
}

func (unit *Unit) GetType() worldcommonmodel.UnitType {
	return unit._type
}
