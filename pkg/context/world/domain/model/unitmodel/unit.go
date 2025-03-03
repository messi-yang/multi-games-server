package unitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

// Unit here is only for reading purpose, for writing units,
// please check the unit model of the type you are looking for.
type Unit struct {
	UnitEntity
}

// Interface Implementation Check
var _ domain.Aggregate = (*Unit)(nil)

func LoadUnit(
	id UnitId,
	worldId globalcommonmodel.WorldId,
	position worldcommonmodel.Position,
	itemId worldcommonmodel.ItemId,
	direction worldcommonmodel.Direction,
	dimension worldcommonmodel.Dimension,
	label *string,
	color *globalcommonmodel.Color,
	_type worldcommonmodel.UnitType,
) Unit {
	return Unit{
		NewUnitEntity(id,
			worldId,
			position,
			itemId,
			direction,
			dimension,
			label,
			color,
			_type,
		),
	}
}

func (unit *Unit) GetId() UnitId {
	return unit.UnitEntity.GetId()
}
