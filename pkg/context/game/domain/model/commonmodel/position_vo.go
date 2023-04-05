package commonmodel

type PositionVo struct {
	x int
	z int
}

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
