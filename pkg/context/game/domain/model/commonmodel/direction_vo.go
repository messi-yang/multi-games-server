package commonmodel

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/valueobject"
)

var (
	ErrNoSuchDirection = errors.New("direction must be one of 0, 1, 2, 3")
)

type DirectionVo int8

// Interface Implementation Check
var _ valueobject.ValueObject[DirectionVo] = (*DirectionVo)(nil)

func NewDirectionVo(direction int8) DirectionVo {
	return DirectionVo(direction % 4)
}

func NewDownDirectionVo() DirectionVo {
	return DirectionVo(0)
}

func (direction DirectionVo) Int8() int8 {
	return int8(direction)
}

func (direction DirectionVo) IsEqual(otherDirection DirectionVo) bool {
	return direction == otherDirection
}

func (direction DirectionVo) IsDown() bool {
	return direction == 0
}

func (direction DirectionVo) IsLeft() bool {
	return direction == 3
}

func (direction DirectionVo) IsUp() bool {
	return direction == 2
}

func (direction DirectionVo) IsRight() bool {
	return direction == 1
}

func (direction DirectionVo) Rotate() DirectionVo {
	return NewDirectionVo(direction.Int8() + 1)
}
