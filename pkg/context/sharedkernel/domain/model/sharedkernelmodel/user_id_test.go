package sharedkernelmodel

import (
	"testing"

	"github.com/google/uuid"
)

func Test_UserId_IsEqual(t *testing.T) {
	uuid1 := uuid.New()
	uuid2 := uuid.New()
	userId1 := NewUserId(uuid1)
	userId2 := NewUserId(uuid1)
	userId3 := NewUserId(uuid2)

	if !userId1.IsEqual(userId2) {
		t.Errorf("userId1 is expected to be equal to userId2")
	}
	if userId1.IsEqual(userId3) {
		t.Errorf("userId1 is expected to be not equal to userId3")
	}
}

func Test_UserId_Uuid(t *testing.T) {
	uuid := uuid.New()
	userId := NewUserId(uuid)

	if userId.Uuid() != uuid {
		t.Errorf("userId should export uuid correctly")
	}
}
