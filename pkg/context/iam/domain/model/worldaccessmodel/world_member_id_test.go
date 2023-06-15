package worldaccessmodel

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewWorldMemberId(t *testing.T) {
	uuid1 := uuid.New()

	worldMemberId1 := NewWorldMemberId(uuid1)
	assert.Equal(t, WorldMemberId{uuid1}, worldMemberId1, "worldMemberId1 should have the provided UUID")
}

func TestWorldMemberId_IsEqual(t *testing.T) {
	uuid1 := uuid.New()
	uuid2 := uuid.New()

	worldMemberId1 := NewWorldMemberId(uuid1)
	worldMemberId2 := NewWorldMemberId(uuid1)
	worldMemberId3 := NewWorldMemberId(uuid2)

	assert.True(t, worldMemberId1.IsEqual(worldMemberId2), "worldMemberId1 should be equal to worldMemberId2")
	assert.False(t, worldMemberId1.IsEqual(worldMemberId3), "worldMemberId1 should not be equal to worldMemberId3")
}

func TestWorldMemberId_Uuid(t *testing.T) {
	uuid := uuid.New()
	worldMemberId := NewWorldMemberId(uuid)

	assert.Equal(t, uuid, worldMemberId.Uuid(), "Uuid() should return the UUID of the worldMemberId")
}
