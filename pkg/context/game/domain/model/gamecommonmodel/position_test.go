package gamecommonmodel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPosition(t *testing.T) {
	position := NewPosition(3, 5)
	assert.Equal(t, Position{3, 5}, position, "position should have x=3 and z=5")
}

func TestPosition_IsEqual(t *testing.T) {
	position1 := NewPosition(3, 5)
	position2 := NewPosition(3, 5)
	position3 := NewPosition(1, 1)

	assert.True(t, position1.IsEqual(position2), "position1 should be equal to position2")
	assert.False(t, position1.IsEqual(position3), "position1 should not be equal to position3")
}

func TestPosition_GetX(t *testing.T) {
	position := NewPosition(3, 5)
	assert.Equal(t, 3, position.GetX(), "GetX() should return 3")
}

func TestPosition_GetZ(t *testing.T) {
	position := NewPosition(3, 5)
	assert.Equal(t, 5, position.GetZ(), "GetZ() should return 5")
}

func TestPosition_Shift(t *testing.T) {
	position := NewPosition(3, 5)
	shiftedPosition := position.Shift(2, -1)
	assert.Equal(t, Position{5, 4}, shiftedPosition, "shifted position should have x=5 and z=4")
}
