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

func (c LocationVo) HasNegativeAxis() bool {
	return c.x < 0 || c.y < 0
}

func (c LocationVo) Shift(x int, y int) LocationVo {
	return NewLocationVo(c.x+x, c.y+y)
}
