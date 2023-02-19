package commonmodel

type LocationVo struct {
	x int
	z int
}

func NewLocationVo(x int, z int) LocationVo {
	return LocationVo{
		x: x,
		z: z,
	}
}

func (vo LocationVo) IsEqual(location LocationVo) bool {
	return vo.x == location.x && vo.z == location.z
}

func (vo LocationVo) GetX() int {
	return vo.x
}

func (vo LocationVo) GetZ() int {
	return vo.z
}

func (vo LocationVo) HasNegativeAxis() bool {
	return vo.x < 0 || vo.z < 0
}

func (vo LocationVo) shift(x int, z int) LocationVo {
	return NewLocationVo(vo.x+x, vo.z+z)
}

func (vo LocationVo) MoveToward(direction DirectionVo, distance int) LocationVo {
	var newLocation LocationVo = vo
	if direction.IsUp() {
		newLocation = newLocation.shift(0, -1*distance)
	} else if direction.IsRight() {
		newLocation = newLocation.shift(1*distance, 0)
	} else if direction.IsDown() {
		newLocation = newLocation.shift(0, 1*distance)
	} else if direction.IsLeft() {
		newLocation = newLocation.shift(-1*distance, 0)
	}
	return newLocation
}
