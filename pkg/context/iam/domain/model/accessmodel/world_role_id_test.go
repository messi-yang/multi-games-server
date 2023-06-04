package accessmodel

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewWorldRoleId(t *testing.T) {
	uuid1 := uuid.New()

	worldRoleId1 := NewWorldRoleId(uuid1)
	assert.Equal(t, WorldRoleId{uuid1}, worldRoleId1, "worldRoleId1 should have the provided UUID")
}

func TestWorldRoleId_IsEqual(t *testing.T) {
	uuid1 := uuid.New()
	uuid2 := uuid.New()

	worldRoleId1 := NewWorldRoleId(uuid1)
	worldRoleId2 := NewWorldRoleId(uuid1)
	worldRoleId3 := NewWorldRoleId(uuid2)

	assert.True(t, worldRoleId1.IsEqual(worldRoleId2), "worldRoleId1 should be equal to worldRoleId2")
	assert.False(t, worldRoleId1.IsEqual(worldRoleId3), "worldRoleId1 should not be equal to worldRoleId3")
}

func TestWorldRoleId_Uuid(t *testing.T) {
	uuid := uuid.New()
	worldRoleId := NewWorldRoleId(uuid)

	assert.Equal(t, uuid, worldRoleId.Uuid(), "Uuid() should return the UUID of the worldRoleId")
}
