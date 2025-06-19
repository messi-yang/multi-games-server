package roomaccessmodel

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewRoomMemberId(t *testing.T) {
	uuid1 := uuid.New()

	roomMemberId1 := NewRoomMemberId(uuid1)
	assert.Equal(t, RoomMemberId{uuid1}, roomMemberId1, "roomMemberId1 should have the provided UUID")
}

func TestRoomMemberId_IsEqual(t *testing.T) {
	uuid1 := uuid.New()
	uuid2 := uuid.New()

	roomMemberId1 := NewRoomMemberId(uuid1)
	roomMemberId2 := NewRoomMemberId(uuid1)
	roomMemberId3 := NewRoomMemberId(uuid2)

	assert.True(t, roomMemberId1.IsEqual(roomMemberId2), "roomMemberId1 should be equal to roomMemberId2")
	assert.False(t, roomMemberId1.IsEqual(roomMemberId3), "roomMemberId1 should not be equal to roomMemberId3")
}

func TestRoomMemberId_Uuid(t *testing.T) {
	uuid := uuid.New()
	roomMemberId := NewRoomMemberId(uuid)

	assert.Equal(t, uuid, roomMemberId.Uuid(), "Uuid() should return the UUID of the roomMemberId")
}
