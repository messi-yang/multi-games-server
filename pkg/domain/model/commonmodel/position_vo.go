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

func (vo PositionVo) shift(x int, z int) PositionVo {
	return NewPositionVo(vo.x+x, vo.z+z)
}

func (vo PositionVo) MoveToward(direction DirectionVo, distance int) PositionVo {
	var newPosition PositionVo = vo
	if direction.IsUp() {
		newPosition = newPosition.shift(0, -1*distance)
	} else if direction.IsRight() {
		newPosition = newPosition.shift(1*distance, 0)
	} else if direction.IsDown() {
		newPosition = newPosition.shift(0, 1*distance)
	} else if direction.IsLeft() {
		newPosition = newPosition.shift(-1*distance, 0)
	}
	return newPosition
}
