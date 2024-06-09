package portalunitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type PortalUnit struct {
	unitmodel.UnitEntity
	targetUnitId *PortalUnitId
}

// Interface Implementation Check
var _ domain.Aggregate = (*PortalUnit)(nil)

func NewPortalUnit(
	id PortalUnitId,
	worldId globalcommonmodel.WorldId,
	position worldcommonmodel.Position,
	itemId worldcommonmodel.ItemId,
	direction worldcommonmodel.Direction,
	dimension worldcommonmodel.Dimension,
	targetUnitId *PortalUnitId,
) PortalUnit {
	return PortalUnit{
		UnitEntity: unitmodel.NewUnitEntity(
			unitmodel.NewUnitId(id.Uuid()),
			worldId,
			position,
			itemId,
			direction,
			dimension,
			nil,
			worldcommonmodel.NewPortalUnitType(),
			nil,
		),
		targetUnitId: targetUnitId,
	}
}

func LoadPortalUnit(
	id PortalUnitId,
	worldId globalcommonmodel.WorldId,
	position worldcommonmodel.Position,
	itemId worldcommonmodel.ItemId,
	direction worldcommonmodel.Direction,
	dimension worldcommonmodel.Dimension,
	targetUnitId *PortalUnitId,
) PortalUnit {
	return PortalUnit{
		UnitEntity: unitmodel.LoadUnitEntity(
			unitmodel.NewUnitId(id.Uuid()),
			worldId,
			position,
			itemId,
			direction,
			dimension,
			nil,
			worldcommonmodel.NewPortalUnitType(),
			nil,
		),
		targetUnitId: targetUnitId,
	}
}

func (unit *PortalUnit) GetId() PortalUnitId {
	return NewPortalUnitId(unit.UnitEntity.GetId().Uuid())
}

func (unit *PortalUnit) GetTargetUnitId() *PortalUnitId {
	return unit.targetUnitId
}

func (unit *PortalUnit) UpdateTargetUnitId(targetUnitId *PortalUnitId) {
	unit.targetUnitId = targetUnitId
}

func (unit *PortalUnit) GetInfoSnapshot() PortalUnitInfo {
	if unit.targetUnitId == nil {
		return PortalUnitInfo{
			TargetUnitId: nil,
		}
	} else {
		return PortalUnitInfo{
			TargetUnitId: lo.TernaryF(
				unit.targetUnitId == nil,
				func() *uuid.UUID { return nil },
				func() *uuid.UUID { return commonutil.ToPointer(unit.targetUnitId.Uuid()) },
			),
		}
	}
}

func (unit *PortalUnit) Delete() {
}
