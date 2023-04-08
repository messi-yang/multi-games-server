package commonmodel

import (
	"testing"
)

func Test_PositionVo_IsEqual(t *testing.T) {
	pos1 := NewPositionVo(0, 0)
	pos2 := NewPositionVo(0, 0)
	pos3 := NewPositionVo(1, 1)

	if !pos1.IsEqual(pos2) {
		t.Errorf("pos1 is expected to be equal to pos2")
	}
	if pos1.IsEqual(pos3) {
		t.Errorf("pos1 is expected to be not equal to pos3")
	}
}
