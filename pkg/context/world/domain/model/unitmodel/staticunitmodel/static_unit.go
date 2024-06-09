package staticunitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

type StaticUnit struct {
	unitmodel.UnitEntity
}

// Interface Implementation Check
var _ domain.Aggregate = (*StaticUnit)(nil)

func NewStaticUnit(
	id StaticUnitId,
	worldId globalcommonmodel.WorldId,
	position worldcommonmodel.Position,
	itemId worldcommonmodel.ItemId,
	direction worldcommonmodel.Direction,
	dimension worldcommonmodel.Dimension,
) StaticUnit {
	return StaticUnit{
		unitmodel.NewUnitEntity(
			unitmodel.NewUnitId(id.Uuid()),
			worldId,
			position,
			itemId,
			direction,
			dimension,
			nil,
			worldcommonmodel.NewStaticUnitType(),
			nil,
		),
	}
}

func LoadStaticUnit(
	id StaticUnitId,
	worldId globalcommonmodel.WorldId,
	position worldcommonmodel.Position,
	itemId worldcommonmodel.ItemId,
	direction worldcommonmodel.Direction,
	dimension worldcommonmodel.Dimension,
) StaticUnit {
	return StaticUnit{
		unitmodel.LoadUnitEntity(
			unitmodel.NewUnitId(id.Uuid()),
			worldId,
			position,
			itemId,
			direction,
			dimension,
			nil,
			worldcommonmodel.NewStaticUnitType(),
			nil,
		),
	}
}

func (unit *StaticUnit) GetId() StaticUnitId {
	return NewStaticUnitId(unit.UnitEntity.GetId().Uuid())
}

func (unit *StaticUnit) Delete() {
}
