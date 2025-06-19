package globalcommonmodel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoomRole(t *testing.T) {
	t.Run("NewRoomRole", func(t *testing.T) {
		roomRole, err := NewRoomRole("admin")
		assert.NoError(t, err)
		assert.Equal(t, roomRole.String(), "admin")

		_, err = NewRoomRole("invalid")
		assert.Error(t, err)
	})

	t.Run("RoomRole", func(t *testing.T) {
		t.Run("IsEqual", func(t *testing.T) {
			roomRole1, _ := NewRoomRole("admin")
			roomRole2, _ := NewRoomRole("admin")
			roomRole3, _ := NewRoomRole("owner")
			assert.True(t, roomRole1.IsEqual(roomRole2))
			assert.False(t, roomRole1.IsEqual(roomRole3))
		})
		t.Run("String", func(t *testing.T) {
			roomRole, _ := NewRoomRole("admin")
			assert.Equal(t, roomRole.String(), "admin")
		})

		t.Run("IsOwner", func(t *testing.T) {
			roomRole, _ := NewRoomRole("owner")
			assert.True(t, roomRole.IsOwner())
		})

		t.Run("IsAdmin", func(t *testing.T) {
			roomRole, _ := NewRoomRole("admin")
			assert.True(t, roomRole.IsAdmin())
		})

		t.Run("IsEditor", func(t *testing.T) {
			roomRole, _ := NewRoomRole("editor")
			assert.True(t, roomRole.IsEditor())
		})

		t.Run("IsViewer", func(t *testing.T) {
			roomRole, _ := NewRoomRole("viewer")
			assert.True(t, roomRole.IsViewer())
		})
	})
}
