package signunitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

type SignUnit struct {
	unitmodel.UnitEntity
}

// Interface Implementation Check
var _ domain.Aggregate = (*SignUnit)(nil)

func NewSignUnit(
	id SignUnitId,
	worldId globalcommonmodel.WorldId,
	position worldcommonmodel.Position,
	itemId worldcommonmodel.ItemId,
	direction worldcommonmodel.Direction,
	dimension worldcommonmodel.Dimension,
	label string,
) SignUnit {
	return SignUnit{
		UnitEntity: unitmodel.NewUnitEntity(
			unitmodel.NewUnitId(id.Uuid()),
			worldId,
			position,
			itemId,
			direction,
			dimension,
			&label,
			nil,
			worldcommonmodel.NewSignUnitType(),
		),
	}
}

func LoadSignUnit(
	id SignUnitId,
	worldId globalcommonmodel.WorldId,
	position worldcommonmodel.Position,
	itemId worldcommonmodel.ItemId,
	direction worldcommonmodel.Direction,
	dimension worldcommonmodel.Dimension,
	label string,
) SignUnit {
	return SignUnit{
		UnitEntity: unitmodel.LoadUnitEntity(
			unitmodel.NewUnitId(id.Uuid()),
			worldId,
			position,
			itemId,
			direction,
			dimension,
			&label,
			nil,
			worldcommonmodel.NewSignUnitType(),
		),
	}
}

func (unit *SignUnit) GetId() SignUnitId {
	return NewSignUnitId(unit.UnitEntity.GetId().Uuid())
}

func (unit *SignUnit) GetLabel() string {
	return *unit.UnitEntity.GetLabel()
}

func (unit *SignUnit) Delete() {
}
