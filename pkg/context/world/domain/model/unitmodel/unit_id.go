package unitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

type UnitId struct {
	worldId  globalcommonmodel.WorldId
	position worldcommonmodel.Position
}

// Interface Implementation Check
var _ domain.ValueObject[UnitId] = (*UnitId)(nil)

func NewUnitId(worldId globalcommonmodel.WorldId, position worldcommonmodel.Position) UnitId {
	return UnitId{
		worldId:  worldId,
		position: position,
	}
}

func (unitId UnitId) IsEqual(otherUnitId UnitId) bool {
	return unitId.worldId.IsEqual(otherUnitId.worldId) && unitId.position.IsEqual(otherUnitId.position)
}

func (unitId UnitId) GetWorldId() globalcommonmodel.WorldId {
	return unitId.worldId
}

func (unitId UnitId) GetPosition() worldcommonmodel.Position {
	return unitId.position
}
