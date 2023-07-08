package commonmodel

import "github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"

type Position struct {
	x int
	z int
}

// Interface Implementation Check
var _ domain.ValueObject[Position] = (*Position)(nil)

func NewPosition(x int, z int) Position {
	return Position{
		x: x,
		z: z,
	}
}

func (position Position) IsEqual(otherPosition Position) bool {
	return position.x == otherPosition.x && position.z == otherPosition.z
}

func (position Position) GetX() int {
	return position.x
}

func (position Position) GetZ() int {
	return position.z
}

func (position Position) Shift(x int, z int) Position {
	return NewPosition(position.x+x, position.z+z)
}
