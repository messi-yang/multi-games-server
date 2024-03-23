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
		t.Run("Should rotate by 90deg in counterclockwise", func(t *testing.T) {
			dimension, _ := worldcommonmodel.NewDimension(1, 1)
			downDirection := worldcommonmodel.NewDownDirection()

			unitEntity := NewUnitEntity(
				NewUnitId(uuid.New()),
				globalcommonmodel.NewWorldId(uuid.New()),
				worldcommonmodel.NewPosition(0, 0),
				worldcommonmodel.NewItemId(uuid.New()),
				downDirection,
				dimension,
				nil,
				worldcommonmodel.NewStaticUnitType(),
				nil,
			)

			rightDirection := worldcommonmodel.NewRightDirection()
			unitEntity.Rotate()
			assert.True(t, unitEntity.GetDirection().IsEqual(rightDirection))
		})
	})

	// t.Run("GetOccupiedPositions", func(t *testing.T) {

	// })
}
