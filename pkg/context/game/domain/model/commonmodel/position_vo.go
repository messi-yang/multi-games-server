package commonmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/valueobject"

type PositionVo struct {
	x int
	z int
}

// Interface Implementation Check
var _ valueobject.ValueObject[PositionVo] = (*PositionVo)(nil)

func NewPositionVo(x int, z int) PositionVo {
	return PositionVo{
		x: x,
		z: z,
	}
}

func (vo PositionVo) IsEqual(position PositionVo) bool {
	return vo.x == position.x && vo.z == position.z
}

func (vo PositionVo) GetX() int {
	return vo.x
}

func (vo PositionVo) GetZ() int {
	return vo.z
}

func (vo PositionVo) Shift(x int, z int) PositionVo {
	return NewPositionVo(vo.x+x, vo.z+z)
}
