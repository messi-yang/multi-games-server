package globalcommonmodel

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRoomId(t *testing.T) {
	t.Run("NewRoomId", func(t *testing.T) {
		uuid := uuid.New()
		roomId := NewRoomId(uuid)

		assert.Equal(t, uuid, roomId.Uuid(), "created room ID should have the correct UUID")
	})

	t.Run("RoomId", func(t *testing.T) {
		t.Run("IsEqual", func(t *testing.T) {
			uuid1 := uuid.New()
			uuid2 := uuid.New()
			roomId1 := NewRoomId(uuid1)
			roomId2 := NewRoomId(uuid1)
			roomId3 := NewRoomId(uuid2)

			assert.True(t, roomId1.IsEqual(roomId2), "room ID 1 should be equal to room ID 2")
			assert.False(t, roomId1.IsEqual(roomId3), "room ID 1 should not be equal to room ID 3")
		})
	})
}
