package commonmodel

import (
	"testing"
)

func Test_Position_IsEqual(t *testing.T) {
	pos1 := NewPosition(0, 0)
	pos2 := NewPosition(0, 0)
	pos3 := NewPosition(1, 1)

	if !pos1.IsEqual(pos2) {
		t.Errorf("pos1 is expected to be equal to pos2")
	}
	if pos1.IsEqual(pos3) {
		t.Errorf("pos1 is expected to be not equal to pos3")
	}
}
