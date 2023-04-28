package commonmodel

import (
	"testing"

	"github.com/google/uuid"
)

func Test_ItemId_IsEqual(t *testing.T) {
	uuid1 := uuid.New()
	uuid2 := uuid.New()
	itemId1 := NewItemId(uuid1)
	itemId2 := NewItemId(uuid1)
	itemId3 := NewItemId(uuid2)

	if !itemId1.IsEqual(itemId2) {
		t.Errorf("itemId1 is expected to be equal to itemId2")
	}
	if itemId1.IsEqual(itemId3) {
		t.Errorf("itemId1 is expected to be not equal to itemId3")
	}
}

func Test_ItemId_Uuid(t *testing.T) {
	uuid := uuid.New()
	itemId := NewItemId(uuid)

	if itemId.Uuid() != uuid {
		t.Errorf("itemId should export uuid correctly")
	}
}
