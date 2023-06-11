package worldaccessmodel

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewUserWorldRoleId(t *testing.T) {
	uuid1 := uuid.New()

	userWorldRoleId1 := NewUserWorldRoleId(uuid1)
	assert.Equal(t, UserWorldRoleId{uuid1}, userWorldRoleId1, "userWorldRoleId1 should have the provided UUID")
}

func TestUserWorldRoleId_IsEqual(t *testing.T) {
	uuid1 := uuid.New()
	uuid2 := uuid.New()

	userWorldRoleId1 := NewUserWorldRoleId(uuid1)
	userWorldRoleId2 := NewUserWorldRoleId(uuid1)
	userWorldRoleId3 := NewUserWorldRoleId(uuid2)

	assert.True(t, userWorldRoleId1.IsEqual(userWorldRoleId2), "userWorldRoleId1 should be equal to userWorldRoleId2")
	assert.False(t, userWorldRoleId1.IsEqual(userWorldRoleId3), "userWorldRoleId1 should not be equal to userWorldRoleId3")
}

func TestUserWorldRoleId_Uuid(t *testing.T) {
	uuid := uuid.New()
	userWorldRoleId := NewUserWorldRoleId(uuid)

	assert.Equal(t, uuid, userWorldRoleId.Uuid(), "Uuid() should return the UUID of the userWorldRoleId")
}
