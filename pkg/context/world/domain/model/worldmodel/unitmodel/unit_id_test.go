package unitmodel

import (
	"testing"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewUnitId(t *testing.T) {
	worldId := globalcommonmodel.NewWorldId(uuid.New())
	position := worldcommonmodel.NewPosition(10, 20)

	unitId := NewUnitId(worldId, position)

	expectedUnitId := UnitId{
		worldId:  worldId,
		position: position,
	}
	assert.Equal(t, expectedUnitId, unitId, "unitId should have the provided worldId and position")
}

func TestUnitId_IsEqual(t *testing.T) {
	worldId1 := globalcommonmodel.NewWorldId(uuid.New())
	worldId2 := globalcommonmodel.NewWorldId(uuid.New())
	position1 := worldcommonmodel.NewPosition(10, 20)
	position2 := worldcommonmodel.NewPosition(30, 40)

	unitId1 := NewUnitId(worldId1, position1)
	unitId2 := NewUnitId(worldId1, position1)
	unitId3 := NewUnitId(worldId2, position1)
	unitId4 := NewUnitId(worldId1, position2)

	assert.True(t, unitId1.IsEqual(unitId2), "unitId1 should be equal to unitId2")
	assert.False(t, unitId1.IsEqual(unitId3), "unitId1 should not be equal to unitId3")
	assert.False(t, unitId1.IsEqual(unitId4), "unitId1 should not be equal to unitId4")
}

func TestUnitId_GetWorldId(t *testing.T) {
	worldId := globalcommonmodel.NewWorldId(uuid.New())
	position := worldcommonmodel.NewPosition(10, 20)
	unitId := NewUnitId(worldId, position)

	assert.Equal(t, worldId, unitId.GetWorldId(), "GetWorldId() should return the worldId of the unitId")
}

func TestUnitId_GetPosition(t *testing.T) {
	worldId := globalcommonmodel.NewWorldId(uuid.New())
	position := worldcommonmodel.NewPosition(10, 20)
	unitId := NewUnitId(worldId, position)

	assert.Equal(t, position, unitId.GetPosition(), "GetPosition() should return the position of the unitId")
}
