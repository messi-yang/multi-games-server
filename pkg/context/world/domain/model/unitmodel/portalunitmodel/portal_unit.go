package portalunitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

type PortalUnit struct {
	unitmodel.UnitEntity
	targetPosition *worldcommonmodel.Position
}

// Interface Implementation Check
var _ domain.Aggregate[PortalUnitId] = (*PortalUnit)(nil)

func NewPortalUnit(
	id PortalUnitId,
	worldId globalcommonmodel.WorldId,
	position worldcommonmodel.Position,
	itemId worldcommonmodel.ItemId,
	direction worldcommonmodel.Direction,
	targetPosition *worldcommonmodel.Position,
) PortalUnit {
	return PortalUnit{
		UnitEntity: unitmodel.NewUnitEntity(
			unitmodel.NewUnitId(id.Uuid()),
			worldId,
			position,
			itemId,
			direction,
			nil,
			worldcommonmodel.NewPortalUnitType(),
			nil,
		),
		targetPosition: targetPosition,
	}
}

func LoadPortalUnit(
	id PortalUnitId,
	worldId globalcommonmodel.WorldId,
	position worldcommonmodel.Position,
	itemId worldcommonmodel.ItemId,
	direction worldcommonmodel.Direction,
	targetPosition *worldcommonmodel.Position,
) PortalUnit {
	return PortalUnit{
		UnitEntity: unitmodel.LoadUnitEntity(
			unitmodel.NewUnitId(id.Uuid()),
			worldId,
			position,
			itemId,
			direction,
			nil,
			worldcommonmodel.NewPortalUnitType(),
			nil,
		),
		targetPosition: targetPosition,
	}
}

func (unit *PortalUnit) GetId() PortalUnitId {
	return NewPortalUnitId(unit.UnitEntity.GetId().Uuid())
}

func (unit *PortalUnit) GetTargetPosition() *worldcommonmodel.Position {
	return unit.targetPosition
}

func (unit *PortalUnit) UpdateTargetPosition(targetPosition *worldcommonmodel.Position) {
	unit.targetPosition = targetPosition
}

func (unit *PortalUnit) GetInfoSnapshot() PortalUnitInfo {
	if unit.targetPosition == nil {
		return PortalUnitInfo{
			TargetPos: nil,
		}
	} else {
		return PortalUnitInfo{
			TargetPos: &struct {
				X int `json:"x"`
				Z int `json:"z"`
			}{
				X: unit.targetPosition.GetX(),
				Z: unit.targetPosition.GetZ(),
			},
		}
	}
}

func (unit *PortalUnit) Delete() {
}
