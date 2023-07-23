package globalcommonmodel

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestWorldId(t *testing.T) {
	t.Run("NewWorldId", func(t *testing.T) {
		uuid := uuid.New()
		worldId := NewWorldId(uuid)

		assert.Equal(t, uuid, worldId.Uuid(), "created world ID should have the correct UUID")
	})

	t.Run("WorldId", func(t *testing.T) {
		t.Run("IsEqual", func(t *testing.T) {
			uuid1 := uuid.New()
			uuid2 := uuid.New()
			worldId1 := NewWorldId(uuid1)
			worldId2 := NewWorldId(uuid1)
			worldId3 := NewWorldId(uuid2)

			assert.True(t, worldId1.IsEqual(worldId2), "world ID 1 should be equal to world ID 2")
			assert.False(t, worldId1.IsEqual(worldId3), "world ID 1 should not be equal to world ID 3")
		})
	})
}
