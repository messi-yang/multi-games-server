package gameaccountmodel

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewGameAccountId(t *testing.T) {
	uuid1 := uuid.New()

	gameAccountId1 := NewGameAccountId(uuid1)
	assert.Equal(t, GameAccountId{uuid1}, gameAccountId1, "gameAccountId1 should have the provided UUID")
}

func TestGameAccountId_IsEqual(t *testing.T) {
	uuid1 := uuid.New()
	uuid2 := uuid.New()

	gameAccountId1 := NewGameAccountId(uuid1)
	gameAccountId2 := NewGameAccountId(uuid1)
	gameAccountId3 := NewGameAccountId(uuid2)

	assert.True(t, gameAccountId1.IsEqual(gameAccountId2), "gameAccountId1 should be equal to gameAccountId2")
	assert.False(t, gameAccountId1.IsEqual(gameAccountId3), "gameAccountId1 should not be equal to gameAccountId3")
}

func TestGameAccountId_Uuid(t *testing.T) {
	uuid := uuid.New()
	gameAccountId := NewGameAccountId(uuid)

	assert.Equal(t, uuid, gameAccountId.Uuid(), "Uuid() should return the UUID of the gameAccountId")
}
