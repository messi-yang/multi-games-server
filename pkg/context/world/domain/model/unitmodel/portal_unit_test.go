package unitmodel

import (
	"testing"

	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPortalUnit(t *testing.T) {
	t.Run("NewPortalUnit", func(t *testing.T) {
		targetPosition := worldcommonmodel.NewPosition(0, 0)
		portalUnit := NewPortalUnit(&targetPosition)
		assert.Equal(t, portalUnit.GetTargetPosition(), targetPosition)
	})

	t.Run("LoadPortalUnit", func(t *testing.T) {
		portalUnitId := NewPortalUnitId(uuid.New())
		targetPosition := worldcommonmodel.NewPosition(0, 0)
		portalUnit := LoadPortalUnit(portalUnitId, &targetPosition)
		assert.True(t, portalUnit.GetId().IsEqual(portalUnitId))
	})

	t.Run("PortalUnit", func(t *testing.T) {
		t.Run("GetTargetPosition", func(t *testing.T) {
			targetPosition := worldcommonmodel.NewPosition(0, 0)
			portalUnit := NewPortalUnit(&targetPosition)
			assert.Equal(t, portalUnit.GetTargetPosition(), targetPosition)
		})
	})
}
