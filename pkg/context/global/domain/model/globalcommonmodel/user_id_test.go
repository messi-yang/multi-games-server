package globalcommonmodel

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserId(t *testing.T) {
	t.Run("NewUserId", func(t *testing.T) {
		uuid := uuid.New()
		userId := NewUserId(uuid)

		assert.Equal(t, uuid, userId.Uuid(), "created user ID should have the correct UUID")
	})

	t.Run("UserId", func(t *testing.T) {
		t.Run("IsEqual", func(t *testing.T) {
			uuid1 := uuid.New()
			uuid2 := uuid.New()
			userId1 := NewUserId(uuid1)
			userId2 := NewUserId(uuid1)
			userId3 := NewUserId(uuid2)

			assert.True(t, userId1.IsEqual(userId2), "user ID 1 should be equal to user ID 2")
			assert.False(t, userId1.IsEqual(userId3), "user ID 1 should not be equal to user ID 3")
		})
	})
}
