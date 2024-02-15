package portalunitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

type PortalUnit struct {
	id             PortalUnitId
	worldId        globalcommonmodel.WorldId
	position       worldcommonmodel.Position
	itemId         worldcommonmodel.ItemId
	direction      worldcommonmodel.Direction
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
		id:             id,
		worldId:        worldId,
		position:       position,
		itemId:         itemId,
		direction:      direction,
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
		id:             id,
		worldId:        worldId,
		position:       position,
		itemId:         itemId,
		direction:      direction,
		targetPosition: targetPosition,
	}
}

func (portalUnit *PortalUnit) GetId() PortalUnitId {
	return portalUnit.id
}

func (portalUnit *PortalUnit) GetWorldId() globalcommonmodel.WorldId {
	return portalUnit.worldId
}

func (portalUnit *PortalUnit) GetPosition() worldcommonmodel.Position {
	return portalUnit.position
}

func (portalUnit *PortalUnit) GetItemId() worldcommonmodel.ItemId {
	return portalUnit.itemId
}

func (portalUnit *PortalUnit) GetDirection() worldcommonmodel.Direction {
	return portalUnit.direction
}

func (portalUnit *PortalUnit) GetTargetPosition() *worldcommonmodel.Position {
	return portalUnit.targetPosition
}

func (portalUnit *PortalUnit) UpdateTargetPosition(targetPosition *worldcommonmodel.Position) {
	portalUnit.targetPosition = targetPosition
}

func (portalUnit *PortalUnit) Rotate() {
	portalUnit.direction = portalUnit.direction.Rotate()
}

func (portalUnit *PortalUnit) GetInfoSnapshot() PortalUnitInfo {
	if portalUnit.targetPosition == nil {
		return PortalUnitInfo{
			TargetPos: nil,
		}
	} else {
		return PortalUnitInfo{
			TargetPos: &struct {
				X int `json:"x"`
				Z int `json:"z"`
			}{
				X: portalUnit.targetPosition.GetX(),
				Z: portalUnit.targetPosition.GetZ(),
			},
		}
	}
}

func (portalUnit *PortalUnit) Delete() {
}
