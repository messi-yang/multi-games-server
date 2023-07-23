package worldcommonmodel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDirection(t *testing.T) {
	direction := NewDirection(3)
	assert.Equal(t, Direction(3), direction, "direction should be 3")

	direction = NewDirection(8)
	assert.Equal(t, Direction(0), direction, "direction should be 0")

	direction = NewDirection(-2)
	assert.Equal(t, Direction(2), direction, "direction should be 2")
}

func TestNewDownDirection(t *testing.T) {
	direction := NewDownDirection()
	assert.Equal(t, Direction(0), direction, "direction should be 0")
}

func TestDirection_Int8(t *testing.T) {
	direction := NewDirection(2)
	int8Value := direction.Int8()
	assert.Equal(t, int8(2), int8Value, "Int8() should return 2")
}

func TestDirection_IsEqual(t *testing.T) {
	direction1 := NewDirection(1)
	direction2 := NewDirection(1)
	direction3 := NewDirection(2)

	assert.True(t, direction1.IsEqual(direction2), "direction1 should be equal to direction2")
	assert.False(t, direction1.IsEqual(direction3), "direction1 should not be equal to direction3")
}

func TestDirection_IsDown(t *testing.T) {
	direction := NewDirection(0)
	assert.True(t, direction.IsDown(), "direction should be down")

	direction = NewDirection(1)
	assert.False(t, direction.IsDown(), "direction should not be down")
}

func TestDirection_IsLeft(t *testing.T) {
	direction := NewDirection(3)
	assert.True(t, direction.IsLeft(), "direction should be left")

	direction = NewDirection(1)
	assert.False(t, direction.IsLeft(), "direction should not be left")
}

func TestDirection_IsUp(t *testing.T) {
	direction := NewDirection(2)
	assert.True(t, direction.IsUp(), "direction should be up")

	direction = NewDirection(0)
	assert.False(t, direction.IsUp(), "direction should not be up")
}

func TestDirection_IsRight(t *testing.T) {
	direction := NewDirection(1)
	assert.True(t, direction.IsRight(), "direction should be right")

	direction = NewDirection(3)
	assert.False(t, direction.IsRight(), "direction should not be right")
}

func TestDirection_Rotate(t *testing.T) {
	direction := NewDirection(2)
	rotatedDirection := direction.Rotate()
	assert.Equal(t, Direction(3), rotatedDirection, "rotated direction should be 3")

	direction = NewDirection(3)
	rotatedDirection = direction.Rotate()
	assert.Equal(t, Direction(0), rotatedDirection, "rotated direction should be 0")
}
