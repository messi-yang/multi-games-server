package fenceunitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

type FenceUnit struct {
	unitmodel.UnitEntity
}

// Interface Implementation Check
var _ domain.Aggregate[FenceUnitId] = (*FenceUnit)(nil)

func NewFenceUnit(
	id FenceUnitId,
	worldId globalcommonmodel.WorldId,
	position worldcommonmodel.Position,
	itemId worldcommonmodel.ItemId,
	direction worldcommonmodel.Direction,
	dimension worldcommonmodel.Dimension,
) FenceUnit {
	return FenceUnit{
		UnitEntity: unitmodel.NewUnitEntity(
			unitmodel.NewUnitId(id.Uuid()),
			worldId,
			position,
			itemId,
			direction,
			dimension,
			nil,
			worldcommonmodel.NewFenceUnitType(),
			nil,
		),
	}
}

func LoadFenceUnit(
	id FenceUnitId,
	worldId globalcommonmodel.WorldId,
	position worldcommonmodel.Position,
	itemId worldcommonmodel.ItemId,
	direction worldcommonmodel.Direction,
	dimension worldcommonmodel.Dimension,
) FenceUnit {
	return FenceUnit{
		UnitEntity: unitmodel.LoadUnitEntity(
			unitmodel.NewUnitId(id.Uuid()),
			worldId,
			position,
			itemId,
			direction,
			dimension,
			nil,
			worldcommonmodel.NewFenceUnitType(),
			nil,
		),
	}
}

func (unit *FenceUnit) GetId() FenceUnitId {
	return NewFenceUnitId(unit.UnitEntity.GetId().Uuid())
}

func (unit *FenceUnit) Delete() {
}
