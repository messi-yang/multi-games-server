package livegamemodel

import (
	"errors"

	"github.com/samber/lo"
)

var (
	ErrNoSuchDirection = errors.New("direction must be one of 0, 1, 2, 3")
)

type DirectionVo int8

func NewDirectionVo(direction int8) (DirectionVo, error) {
	found := lo.IndexOf([]int8{0, 1, 2, 3}, direction)
	if found == -1 {
		return 0, ErrNoSuchDirection
	}

	return DirectionVo(direction), nil
}

func (direction DirectionVo) ToInt8() int8 {
	return int8(direction)
}

func (direction DirectionVo) IsUp() bool {
	return direction == 0
}

func (direction DirectionVo) IsRight() bool {
	return direction == 1
}

func (direction DirectionVo) IsDown() bool {
	return direction == 2
}

func (direction DirectionVo) IsLeft() bool {
	return direction == 3
}
