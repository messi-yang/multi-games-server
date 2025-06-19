package playermodel

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewPlayerId(t *testing.T) {
	uuid1 := uuid.New()

	playerId1 := NewPlayerId(uuid1)
	assert.Equal(t, PlayerId{uuid1}, playerId1, "playerId1 should have the provided UUID")
}

func TestPlayerId_IsEqual(t *testing.T) {
	uuid1 := uuid.New()
	uuid2 := uuid.New()

	playerId1 := NewPlayerId(uuid1)
	playerId2 := NewPlayerId(uuid1)
	playerId3 := NewPlayerId(uuid2)

	assert.True(t, playerId1.IsEqual(playerId2), "playerId1 should be equal to playerId2")
	assert.False(t, playerId1.IsEqual(playerId3), "playerId1 should not be equal to playerId3")
}

func TestPlayerId_Uuid(t *testing.T) {
	uuid := uuid.New()
	playerId := NewPlayerId(uuid)

	assert.Equal(t, uuid, playerId.Uuid(), "Uuid() should return the UUID of the playerId")
}
