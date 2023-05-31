package accessmodel

import (
	"testing"

	"github.com/google/uuid"
)

func Test_WorldRoleId_IsEqual(t *testing.T) {
	uuid1 := uuid.New()
	uuid2 := uuid.New()
	worldRoleId1 := NewWorldRoleId(uuid1)
	worldRoleId2 := NewWorldRoleId(uuid1)
	worldRoleId3 := NewWorldRoleId(uuid2)

	if !worldRoleId1.IsEqual(worldRoleId2) {
		t.Errorf("worldRoleId1 is expected to be equal to worldRoleId2")
	}
	if worldRoleId1.IsEqual(worldRoleId3) {
		t.Errorf("worldRoleId1 is expected to be not equal to worldRoleId3")
	}
}

func Test_WorldRoleId_Uuid(t *testing.T) {
	uuid := uuid.New()
	worldRoleId := NewWorldRoleId(uuid)

	if worldRoleId.Uuid() != uuid {
		t.Errorf("worldRoleId should export uuid correctly")
	}
}
