package unitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type UnitId struct {
	worldId  sharedkernelmodel.WorldId
	position commonmodel.Position
}

// Interface Implementation Check
var _ domain.ValueObject[UnitId] = (*UnitId)(nil)

func NewUnitId(worldId sharedkernelmodel.WorldId, position commonmodel.Position) UnitId {
	return UnitId{
		worldId:  worldId,
		position: position,
	}
}

func (unitId UnitId) IsEqual(otherUnitId UnitId) bool {
	return unitId.worldId.IsEqual(otherUnitId.worldId) && unitId.position.IsEqual(otherUnitId.position)
}

func (unitId UnitId) GetWorldId() sharedkernelmodel.WorldId {
	return unitId.worldId
}

func (unitId UnitId) GetPosition() commonmodel.Position {
	return unitId.position
}
