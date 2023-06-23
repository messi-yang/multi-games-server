package worldaccountmodel

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewWorldAccountId(t *testing.T) {
	uuid1 := uuid.New()

	worldAccountId1 := NewWorldAccountId(uuid1)
	assert.Equal(t, WorldAccountId{uuid1}, worldAccountId1, "worldAccountId1 should have the provided UUID")
}

func TestWorldAccountId_IsEqual(t *testing.T) {
	uuid1 := uuid.New()
	uuid2 := uuid.New()

	worldAccountId1 := NewWorldAccountId(uuid1)
	worldAccountId2 := NewWorldAccountId(uuid1)
	worldAccountId3 := NewWorldAccountId(uuid2)

	assert.True(t, worldAccountId1.IsEqual(worldAccountId2), "worldAccountId1 should be equal to worldAccountId2")
	assert.False(t, worldAccountId1.IsEqual(worldAccountId3), "worldAccountId1 should not be equal to worldAccountId3")
}

func TestWorldAccountId_Uuid(t *testing.T) {
	uuid := uuid.New()
	worldAccountId := NewWorldAccountId(uuid)

	assert.Equal(t, uuid, worldAccountId.Uuid(), "Uuid() should return the UUID of the worldAccountId")
}
