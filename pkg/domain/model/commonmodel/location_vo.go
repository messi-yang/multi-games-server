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

func (c LocationVo) GetX() int {
	return c.x
}

func (c LocationVo) GetZ() int {
	return c.z
}

func (c LocationVo) HasNegativeAxis() bool {
	return c.x < 0 || c.z < 0
}

func (c LocationVo) Shift(x int, z int) LocationVo {
	return NewLocationVo(c.x+x, c.z+z)
}
