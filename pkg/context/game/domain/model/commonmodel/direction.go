package commonmodel

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/domainmodel"
)

var (
	ErrNoSuchDirection = errors.New("direction must be one of 0, 1, 2, 3")
)

type Direction int8

// Interface Implementation Check
var _ domainmodel.ValueObject[Direction] = (*Direction)(nil)

func NewDirection(direction int8) Direction {
	return Direction(direction % 4)
}

func NewDownDirection() Direction {
	return Direction(0)
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
