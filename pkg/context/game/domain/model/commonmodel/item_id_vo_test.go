package commonmodel

import (
	"testing"

	"github.com/google/uuid"
)

func Test_ItemIdVo_IsEqual(t *testing.T) {
	uuid1 := uuid.New()
	uuid2 := uuid.New()
	itemId1 := NewItemIdVo(uuid1)
	itemId2 := NewItemIdVo(uuid1)
	itemId3 := NewItemIdVo(uuid2)

	if !itemId1.IsEqual(itemId2) {
		t.Errorf("zeroValueItemId is expected to be equal to itemId2")
	}
	if itemId1.IsEqual(itemId3) {
		t.Errorf("zeroValueItemId is not expected to be equal to itemId3")
	}
}

func Test_ItemIdVo_Uuid(t *testing.T) {
	uuid := uuid.New()
	itemId := NewItemIdVo(uuid)

	if itemId.Uuid() != uuid {
		t.Errorf("itemId should export uuid correctly")
	}
}
