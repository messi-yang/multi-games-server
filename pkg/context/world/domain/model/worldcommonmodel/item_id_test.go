package worldcommonmodel

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewItemId(t *testing.T) {
	id := uuid.New()
	itemId := NewItemId(id)
	assert.Equal(t, ItemId{id}, itemId, "itemId should have the provided UUID")
}

func TestItemId_IsEqual(t *testing.T) {
	id1 := uuid.New()
	id2 := uuid.New()
	itemId1 := NewItemId(id1)
	itemId2 := NewItemId(id1)
	itemId3 := NewItemId(id2)

	assert.True(t, itemId1.IsEqual(itemId2), "itemId1 should be equal to itemId2")
	assert.False(t, itemId1.IsEqual(itemId3), "itemId1 should not be equal to itemId3")
}

func TestItemId_Uuid(t *testing.T) {
	id := uuid.New()
	itemId := NewItemId(id)

	assert.Equal(t, id, itemId.Uuid(), "itemId should return the UUID correctly")
}
