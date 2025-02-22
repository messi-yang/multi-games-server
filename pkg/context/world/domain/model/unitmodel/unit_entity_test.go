package unitmodel

import (
	"testing"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUnitEntity(t *testing.T) {
	t.Run("Rotate", func(t *testing.T) {
		t.Run("Symmetric unit", func(t *testing.T) {
			t.Run("Should rotate unit by 90 degree", func(t *testing.T) {
				dimension, _ := worldcommonmodel.NewDimension(3, 3)

				unitEntity := NewUnitEntity(
					NewUnitId(uuid.New()),
					globalcommonmodel.NewWorldId(uuid.New()),
					worldcommonmodel.NewPosition(0, 0),
					worldcommonmodel.NewItemId(uuid.New()),
					worldcommonmodel.NewDirection(0),
					dimension,
					nil,
					nil,
					worldcommonmodel.NewStaticUnitType(),
					nil,
				)

				unitEntity.Rotate()
				assert.True(t, unitEntity.GetDirection().IsEqual(worldcommonmodel.NewDirection(1)))

				unitEntity.Rotate()
				assert.True(t, unitEntity.GetDirection().IsEqual(worldcommonmodel.NewDirection(2)))

				unitEntity.Rotate()
				assert.True(t, unitEntity.GetDirection().IsEqual(worldcommonmodel.NewDirection(3)))
			})
		})

		t.Run("Non-symmetric unit", func(t *testing.T) {
			t.Run("Should rotate unit by 180 degree", func(t *testing.T) {
				dimension, _ := worldcommonmodel.NewDimension(2, 3)

				unitEntity1 := NewUnitEntity(
					NewUnitId(uuid.New()),
					globalcommonmodel.NewWorldId(uuid.New()),
					worldcommonmodel.NewPosition(0, 0),
					worldcommonmodel.NewItemId(uuid.New()),
					worldcommonmodel.NewDirection(0),
					dimension,
					nil,
					nil,
					worldcommonmodel.NewStaticUnitType(),
					nil,
				)

				unitEntity1.Rotate()
				assert.True(t, unitEntity1.GetDirection().IsEqual(worldcommonmodel.NewDirection(2)))

				unitEntity1.Rotate()
				assert.True(t, unitEntity1.GetDirection().IsEqual(worldcommonmodel.NewDirection(0)))

				unitEntity2 := NewUnitEntity(
					NewUnitId(uuid.New()),
					globalcommonmodel.NewWorldId(uuid.New()),
					worldcommonmodel.NewPosition(0, 0),
					worldcommonmodel.NewItemId(uuid.New()),
					worldcommonmodel.NewDirection(1),
					dimension,
					nil,
					nil,
					worldcommonmodel.NewStaticUnitType(),
					nil,
				)

				unitEntity2.Rotate()
				assert.True(t, unitEntity2.GetDirection().IsEqual(worldcommonmodel.NewDirection(3)))

				unitEntity2.Rotate()
				assert.True(t, unitEntity2.GetDirection().IsEqual(worldcommonmodel.NewDirection(1)))
			})
		})
	})

	t.Run("GetOccupiedPositions", func(t *testing.T) {
		t.Run("Should have just 1 position and the position is equal to the position of the unit when dimension is 1x1", func(t *testing.T) {
			dimension, _ := worldcommonmodel.NewDimension(1, 1)
			unitPosition := worldcommonmodel.NewPosition(10, 10)

			unitEntity := NewUnitEntity(
				NewUnitId(uuid.New()),
				globalcommonmodel.NewWorldId(uuid.New()),
				unitPosition,
				worldcommonmodel.NewItemId(uuid.New()),
				worldcommonmodel.NewDownDirection(),
				dimension,
				nil,
				nil,
				worldcommonmodel.NewStaticUnitType(),
				nil,
			)

			occupiedPositions := unitEntity.GetOccupiedPositions()
			assert.Equal(t, len(occupiedPositions), 1)
			assert.True(t, occupiedPositions[0].IsEqual(unitPosition))
		})

		t.Run("Should correctly return the occupied positions of the unit with different directions", func(t *testing.T) {
			t.Run("When facing down", func(t *testing.T) {
				dimension, _ := worldcommonmodel.NewDimension(2, 3)
				unitEntity1 := NewUnitEntity(
					NewUnitId(uuid.New()),
					globalcommonmodel.NewWorldId(uuid.New()),
					worldcommonmodel.NewPosition(10, 10),
					worldcommonmodel.NewItemId(uuid.New()),
					worldcommonmodel.NewDownDirection(),
					dimension,
					nil,
					nil,
					worldcommonmodel.NewStaticUnitType(),
					nil,
				)

				occupiedPositions := unitEntity1.GetOccupiedPositions()
				assert.Equal(t, len(occupiedPositions), 6)
				assert.True(t, occupiedPositions[0].IsEqual(worldcommonmodel.NewPosition(10, 10)))
				assert.True(t, occupiedPositions[len(occupiedPositions)-1].IsEqual(worldcommonmodel.NewPosition(11, 12)))
			})

			t.Run("When facing right", func(t *testing.T) {
				dimension, _ := worldcommonmodel.NewDimension(2, 3)
				unitEntity1 := NewUnitEntity(
					NewUnitId(uuid.New()),
					globalcommonmodel.NewWorldId(uuid.New()),
					worldcommonmodel.NewPosition(10, 10),
					worldcommonmodel.NewItemId(uuid.New()),
					worldcommonmodel.NewRightDirection(),
					dimension,
					nil,
					nil,
					worldcommonmodel.NewStaticUnitType(),
					nil,
				)

				occupiedPositions := unitEntity1.GetOccupiedPositions()
				assert.Equal(t, len(occupiedPositions), 6)
				assert.True(t, occupiedPositions[0].IsEqual(worldcommonmodel.NewPosition(10, 10)))
				assert.True(t, occupiedPositions[len(occupiedPositions)-1].IsEqual(worldcommonmodel.NewPosition(12, 11)))
			})

			t.Run("When facing up", func(t *testing.T) {
				dimension, _ := worldcommonmodel.NewDimension(2, 3)
				unitEntity1 := NewUnitEntity(
					NewUnitId(uuid.New()),
					globalcommonmodel.NewWorldId(uuid.New()),
					worldcommonmodel.NewPosition(10, 10),
					worldcommonmodel.NewItemId(uuid.New()),
					worldcommonmodel.NewUpDirection(),
					dimension,
					nil,
					nil,
					worldcommonmodel.NewStaticUnitType(),
					nil,
				)

				occupiedPositions := unitEntity1.GetOccupiedPositions()
				assert.Equal(t, len(occupiedPositions), 6)
				assert.True(t, occupiedPositions[0].IsEqual(worldcommonmodel.NewPosition(10, 10)))
				assert.True(t, occupiedPositions[len(occupiedPositions)-1].IsEqual(worldcommonmodel.NewPosition(11, 12)))
			})

			t.Run("When facing left", func(t *testing.T) {
				dimension, _ := worldcommonmodel.NewDimension(2, 3)
				unitEntity1 := NewUnitEntity(
					NewUnitId(uuid.New()),
					globalcommonmodel.NewWorldId(uuid.New()),
					worldcommonmodel.NewPosition(10, 10),
					worldcommonmodel.NewItemId(uuid.New()),
					worldcommonmodel.NewLeftDirection(),
					dimension,
					nil,
					nil,
					worldcommonmodel.NewStaticUnitType(),
					nil,
				)

				occupiedPositions := unitEntity1.GetOccupiedPositions()
				assert.Equal(t, len(occupiedPositions), 6)
				assert.True(t, occupiedPositions[0].IsEqual(worldcommonmodel.NewPosition(10, 10)))
				assert.True(t, occupiedPositions[len(occupiedPositions)-1].IsEqual(worldcommonmodel.NewPosition(12, 11)))
			})
		})
	})
}
