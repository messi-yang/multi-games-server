package gamermodel

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewGamerId(t *testing.T) {
	uuid1 := uuid.New()

	gamerId1 := NewGamerId(uuid1)
	assert.Equal(t, GamerId{uuid1}, gamerId1, "gamerId1 should have the provided UUID")
}

func TestGamerId_IsEqual(t *testing.T) {
	uuid1 := uuid.New()
	uuid2 := uuid.New()

	gamerId1 := NewGamerId(uuid1)
	gamerId2 := NewGamerId(uuid1)
	gamerId3 := NewGamerId(uuid2)

	assert.True(t, gamerId1.IsEqual(gamerId2), "gamerId1 should be equal to gamerId2")
	assert.False(t, gamerId1.IsEqual(gamerId3), "gamerId1 should not be equal to gamerId3")
}

func TestGamerId_Uuid(t *testing.T) {
	uuid := uuid.New()
	gamerId := NewGamerId(uuid)

	assert.Equal(t, uuid, gamerId.Uuid(), "Uuid() should return the UUID of the gamerId")
}
