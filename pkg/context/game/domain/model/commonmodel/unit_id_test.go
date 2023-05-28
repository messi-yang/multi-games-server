package commonmodel

import (
	"testing"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/google/uuid"
)

func Test_UnitId_IsEqual(t *testing.T) {
	worldId1 := sharedkernelmodel.NewWorldId(uuid.New())
	worldId2 := sharedkernelmodel.NewWorldId(uuid.New())
	pos1 := NewPosition(0, 0)
	pos2 := NewPosition(0, 1)

	unitId1 := NewUnitId(worldId1, pos1)
	unitId2 := NewUnitId(worldId1, pos1)
	unitId3 := NewUnitId(worldId2, pos2)

	if !unitId1.IsEqual(unitId2) {
		t.Errorf("unitId1 is expected to be equal to unitId2")
	}
	if unitId1.IsEqual(unitId3) {
		t.Errorf("unitId1 is expected to be not equal to unitId3")
	}
}
