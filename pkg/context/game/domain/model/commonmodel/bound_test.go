package commonmodel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBound(t *testing.T) {
	pos1 := NewPosition(0, 0)
	pos2 := NewPosition(10, 10)
	_, err := NewBound(pos2, pos1)
	assert.Error(t, err, "should return error when \"from\" exceeds \"to\" in either x or z axis")

	_, err = NewBound(pos1, pos2)
	assert.NoError(t, err, "should get no error when providing valid \"from\" and \"to\" positions")
}

func TestBound_IsEqual(t *testing.T) {
	bound1, _ := NewBound(NewPosition(0, 0), NewPosition(10, 10))
	bound2, _ := NewBound(NewPosition(0, 0), NewPosition(10, 10))
	bound3, _ := NewBound(NewPosition(0, 0), NewPosition(11, 11))

	assert.True(t, bound1.IsEqual(bound2), "bound1 should be equal to bound2")
	assert.False(t, bound1.IsEqual(bound3), "bound1 should not be equal to bound3")
}

func TestBound_GetWidth(t *testing.T) {
	pos1 := NewPosition(0, 0)
	pos2 := NewPosition(10, 10)
	bound, _ := NewBound(pos1, pos2)

	assert.Equal(t, 11, bound.GetWidth(), "bound width should be 11")
}

func TestBound_GetHeight(t *testing.T) {
	pos1 := NewPosition(0, 0)
	pos2 := NewPosition(10, 10)
	bound, _ := NewBound(pos1, pos2)

	assert.Equal(t, 11, bound.GetHeight(), "bound height should be 11")
}

func TestBound_GetCenterPos(t *testing.T) {
	bound1, _ := NewBound(NewPosition(0, 0), NewPosition(10, 10))
	expectedCenterPos := NewPosition(5, 5)
	assert.True(t, bound1.GetCenterPos().IsEqual(expectedCenterPos), "center position of the bound should be (5, 5)")

	bound2, _ := NewBound(NewPosition(0, 0), NewPosition(11, 11))
	assert.True(t, bound2.GetCenterPos().IsEqual(expectedCenterPos), "center position of the bound should be (5, 5)")
}
