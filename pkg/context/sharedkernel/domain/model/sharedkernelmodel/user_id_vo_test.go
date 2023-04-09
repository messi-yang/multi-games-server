package sharedkernelmodel

import (
	"testing"

	"github.com/google/uuid"
)

func Test_UserIdVo_IsEqual(t *testing.T) {
	uuid1 := uuid.New()
	uuid2 := uuid.New()
	userId1 := NewUserIdVo(uuid1)
	userId2 := NewUserIdVo(uuid1)
	userId3 := NewUserIdVo(uuid2)

	if !userId1.IsEqual(userId2) {
		t.Errorf("userId1 is expected to be equal to userId2")
	}
	if userId1.IsEqual(userId3) {
		t.Errorf("userId1 is expected to be not equal to userId3")
	}
}

func Test_UserIdVo_Uuid(t *testing.T) {
	uuid := uuid.New()
	userId := NewUserIdVo(uuid)

	if userId.Uuid() != uuid {
		t.Errorf("userId should export uuid correctly")
	}
}
