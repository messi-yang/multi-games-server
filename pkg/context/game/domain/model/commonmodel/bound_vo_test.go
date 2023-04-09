package commonmodel

import "testing"

func Test_NewBoundVo(t *testing.T) {
	pos1 := NewPositionVo(0, 0)
	pos2 := NewPositionVo(10, 10)
	_, err := NewBoundVo(pos2, pos1)
	if err == nil {
		t.Errorf("should return error when \"from\" exceeds \"to\" in either x or z axis")
	}
	_, err = NewBoundVo(pos1, pos2)
	if err != nil {
		t.Errorf("should get no error when providing valid \"from\" and \"to\" positions")
	}
}

func Test_BoundVo_IsEqual(t *testing.T) {
	bound1, _ := NewBoundVo(NewPositionVo(0, 0), NewPositionVo(10, 10))
	bound2, _ := NewBoundVo(NewPositionVo(0, 0), NewPositionVo(10, 10))
	bound3, _ := NewBoundVo(NewPositionVo(0, 0), NewPositionVo(11, 11))

	if !bound1.IsEqual(bound2) {
		t.Errorf("bound1 should be equal to bound2")
	}
	if bound1.IsEqual(bound3) {
		t.Errorf("bound1 should not be equal to bound3")
	}
}

func Test_BoundVo_GetWidth(t *testing.T) {
	pos1 := NewPositionVo(0, 0)
	pos2 := NewPositionVo(10, 10)
	bound, _ := NewBoundVo(pos1, pos2)

	if bound.GetWidth() != 11 {
		t.Errorf("bound width should be 11")
	}
}

func Test_BoundVo_GetHeight(t *testing.T) {
	pos1 := NewPositionVo(0, 0)
	pos2 := NewPositionVo(10, 10)
	bound, _ := NewBoundVo(pos1, pos2)

	if bound.GetHeight() != 11 {
		t.Errorf("bound height should be 11")
	}
}

func Test_BoundVo_GetCenterPos(t *testing.T) {
	bound1, _ := NewBoundVo(NewPositionVo(0, 0), NewPositionVo(10, 10))
	expectedCenterPos := NewPositionVo(5, 5)
	if !bound1.GetCenterPos().IsEqual(expectedCenterPos) {
		t.Errorf("center position of the bound should be (5, 5)")
	}

	bound2, _ := NewBoundVo(NewPositionVo(0, 0), NewPositionVo(11, 11))
	expectedCenterPos = NewPositionVo(5, 5)
	if !bound2.GetCenterPos().IsEqual(expectedCenterPos) {
		t.Errorf("center position of the bound should be (5, 5)")
	}
}
