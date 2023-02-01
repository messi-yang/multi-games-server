package commonmodel

type LocationVo struct {
	x int
	y int
}

func NewLocationVo(x int, y int) LocationVo {
	return LocationVo{
		x: x,
		y: y,
	}
}

func (c LocationVo) GetX() int {
	return c.x
}

func (c LocationVo) GetY() int {
	return c.y
}
