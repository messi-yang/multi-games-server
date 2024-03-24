package worldcommonmodel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBound(t *testing.T) {
	t.Run("NewBound", func(t *testing.T) {
		pos1 := NewPosition(0, 0)
		pos2 := NewPosition(10, 10)
		_, err := NewBound(pos2, pos1)
		assert.Error(t, err, "should return error when \"from\" exceeds \"to\" in either x or z axis")

		_, err = NewBound(pos1, pos2)
		assert.NoError(t, err, "should get no error when providing valid \"from\" and \"to\" positions")
	})

	t.Run("IsEqual", func(t *testing.T) {
		bound1, _ := NewBound(NewPosition(0, 0), NewPosition(10, 10))
		bound2, _ := NewBound(NewPosition(0, 0), NewPosition(10, 10))
		bound3, _ := NewBound(NewPosition(0, 0), NewPosition(11, 11))

		assert.True(t, bound1.IsEqual(bound2), "bound1 should be equal to bound2")
		assert.False(t, bound1.IsEqual(bound3), "bound1 should not be equal to bound3")
	})

	t.Run("GetRightTop", func(t *testing.T) {
		pos1 := NewPosition(0, 0)
		pos2 := NewPosition(10, 10)
		bound, _ := NewBound(pos1, pos2)

		assert.True(t, bound.GetRightUp().IsEqual(NewPosition(10, 0)))
	})

	t.Run("GetLeftDown", func(t *testing.T) {
		pos1 := NewPosition(0, 0)
		pos2 := NewPosition(10, 10)
		bound, _ := NewBound(pos1, pos2)

		assert.True(t, bound.GetLeftDown().IsEqual(NewPosition(0, 10)))
	})

	t.Run("GetWidth", func(t *testing.T) {
		pos1 := NewPosition(0, 0)
		pos2 := NewPosition(10, 10)
		bound, _ := NewBound(pos1, pos2)

		assert.Equal(t, 11, bound.GetWidth(), "bound width should be 11")
	})

	t.Run("GetHeight", func(t *testing.T) {
		pos1 := NewPosition(0, 0)
		pos2 := NewPosition(10, 10)
		bound, _ := NewBound(pos1, pos2)

		assert.Equal(t, 11, bound.GetHeight(), "bound height should be 11")
	})

	t.Run("Iterate", func(t *testing.T) {
		fromPos := NewPosition(0, 0)
		toPos := NewPosition(10, 10)
		bound, _ := NewBound(fromPos, toPos)

		pos1 := NewPosition(3, 8)
		pos1VisitedCount := 0

		pos2 := NewPosition(9, 2)
		pos2VisitedCount := 0

		pos3 := NewPosition(10, 3)
		pos3VisitedCount := 0

		pos4 := NewPosition(11, 0)
		pos4VisitedCount := 0

		bound.Iterate(func(position Position) {
			if position.IsEqual(pos1) {
				pos1VisitedCount += 1
			}
			if position.IsEqual(pos2) {
				pos2VisitedCount += 1
			}
			if position.IsEqual(pos3) {
				pos3VisitedCount += 1
			}
			if position.IsEqual(pos4) {
				pos4VisitedCount += 1
			}
		})

		assert.Equal(t, pos1VisitedCount, 1)
		assert.Equal(t, pos2VisitedCount, 1)
		assert.Equal(t, pos3VisitedCount, 1)
		assert.Equal(t, pos4VisitedCount, 0)
	})
}
