package linkunitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

type LinkUnit struct {
	id                   LinkUnitId
	worldId              globalcommonmodel.WorldId
	position             worldcommonmodel.Position
	itemId               worldcommonmodel.ItemId
	direction            worldcommonmodel.Direction
	label                *string
	url                  globalcommonmodel.Url
	domainEventCollector *domain.DomainEventCollector
}

// Interface Implementation Check
var _ domain.Aggregate = (*LinkUnit)(nil)

func NewLinkUnit(
	id LinkUnitId,
	worldId globalcommonmodel.WorldId,
	position worldcommonmodel.Position,
	itemId worldcommonmodel.ItemId,
	direction worldcommonmodel.Direction,
	label *string,
	url globalcommonmodel.Url,
) LinkUnit {
	return LinkUnit{
		id:                   id,
		worldId:              worldId,
		position:             position,
		itemId:               itemId,
		direction:            direction,
		label:                label,
		url:                  url,
		domainEventCollector: domain.NewDomainEventCollector(),
	}
}

func LoadLinkUnit(
	id LinkUnitId,
	worldId globalcommonmodel.WorldId,
	position worldcommonmodel.Position,
	itemId worldcommonmodel.ItemId,
	direction worldcommonmodel.Direction,
	label *string,
	url globalcommonmodel.Url,
) LinkUnit {
	return LinkUnit{
		id:                   id,
		worldId:              worldId,
		position:             position,
		itemId:               itemId,
		direction:            direction,
		label:                label,
		url:                  url,
		domainEventCollector: domain.NewDomainEventCollector(),
	}
}

func (unit *LinkUnit) PopDomainEvents() []domain.DomainEvent {
	return unit.domainEventCollector.PopAll()
}

func (unit *LinkUnit) GetId() LinkUnitId {
	return unit.id
}

func (unit *LinkUnit) GetWorldId() globalcommonmodel.WorldId {
	return unit.worldId
}

func (unit *LinkUnit) GetPosition() worldcommonmodel.Position {
	return unit.position
}

func (unit *LinkUnit) GetItemId() worldcommonmodel.ItemId {
	return unit.itemId
}

func (unit *LinkUnit) GetDirection() worldcommonmodel.Direction {
	return unit.direction
}

func (unit *LinkUnit) GetLabel() *string {
	return unit.label
}

func (unit *LinkUnit) GetUrl() globalcommonmodel.Url {
	return unit.url
}

func (unit *LinkUnit) Rotate() {
	unit.direction = unit.direction.Rotate()
}

func (unit *LinkUnit) Delete() {
}
