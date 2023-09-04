package unitmodel

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPortalUnitId(t *testing.T) {
	t.Run("NewPortalUnitId", func(t *testing.T) {
		uuid := uuid.New()
		portalUnitId := NewPortalUnitId(uuid)
		assert.Equal(t, PortalUnitId{uuid}, portalUnitId)
	})

	t.Run("PortalUnitID", func(t *testing.T) {
		t.Run("IsEqual", func(t *testing.T) {
			uuid1 := uuid.New()
			uuid2 := uuid.New()
			portalUnitId1 := NewPortalUnitId(uuid1)
			portalUnitId2 := NewPortalUnitId(uuid1)
			portalUnitId3 := NewPortalUnitId(uuid2)
			assert.True(t, portalUnitId1.IsEqual(portalUnitId2))
			assert.False(t, portalUnitId1.IsEqual(portalUnitId3))
		})
		t.Run("Uuid", func(t *testing.T) {
			uuid := uuid.New()
			portalUnitId := NewPortalUnitId(uuid)
			assert.Equal(t, uuid, portalUnitId.Uuid())
		})
	})
}
