package colorunitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

type ColorUnit struct {
	unitmodel.UnitEntity
	color globalcommonmodel.Color
}

// Interface Implementation Check
var _ domain.Aggregate = (*ColorUnit)(nil)

func NewColorUnit(
	id ColorUnitId,
	worldId globalcommonmodel.WorldId,
	position worldcommonmodel.Position,
	itemId worldcommonmodel.ItemId,
	direction worldcommonmodel.Direction,
	dimension worldcommonmodel.Dimension,
	label *string,
	color *globalcommonmodel.Color,
) ColorUnit {
	return ColorUnit{
		UnitEntity: unitmodel.NewUnitEntity(
			unitmodel.NewUnitId(id.Uuid()),
			worldId,
			position,
			itemId,
			direction,
			dimension,
			label,
			color,
			worldcommonmodel.NewColorUnitType(),
			nil,
		),
		color: *color,
	}
}

func LoadColorUnit(
	id ColorUnitId,
	worldId globalcommonmodel.WorldId,
	position worldcommonmodel.Position,
	itemId worldcommonmodel.ItemId,
	direction worldcommonmodel.Direction,
	dimension worldcommonmodel.Dimension,
	label *string,
	color *globalcommonmodel.Color,
) ColorUnit {
	return ColorUnit{
		UnitEntity: unitmodel.LoadUnitEntity(
			unitmodel.NewUnitId(id.Uuid()),
			worldId,
			position,
			itemId,
			direction,
			dimension,
			label,
			color,
			worldcommonmodel.NewColorUnitType(),
			nil,
		),
		color: *color,
	}
}

func (unit *ColorUnit) GetId() ColorUnitId {
	return NewColorUnitId(unit.UnitEntity.GetId().Uuid())
}

func (unit *ColorUnit) GetColor() globalcommonmodel.Color {
	return unit.color
}

func (unit *ColorUnit) Delete() {
}
