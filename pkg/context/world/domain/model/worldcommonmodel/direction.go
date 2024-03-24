package worldcommonmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
)

type Direction int8

// Interface Implementation Check
var _ domain.ValueObject[Direction] = (*Direction)(nil)

func NewDirection(direction int8) Direction {
	if direction >= 0 {
		return Direction(direction % 4)
	} else {
		return Direction(-direction % 4)
	}
}

func NewDownDirection() Direction {
	return Direction(0)
}

func NewRightDirection() Direction {
	return Direction(1)
}

func NewUpDirection() Direction {
	return Direction(2)
}

func NewLeftDirection() Direction {
	return Direction(3)
}

func (direction Direction) Int8() int8 {
	return int8(direction)
}

func (direction Direction) IsEqual(otherDirection Direction) bool {
	return direction == otherDirection
}

func (direction Direction) IsDown() bool {
	return direction == 0
}

func (direction Direction) IsLeft() bool {
	return direction == 3
}

func (direction Direction) IsUp() bool {
	return direction == 2
}

func (direction Direction) IsRight() bool {
	return direction == 1
}

func (direction Direction) Rotate() Direction {
	return NewDirection(direction.Int8() + 1)
}
