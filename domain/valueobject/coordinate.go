package valueobject

type Coordinate struct {
	x int
	y int
}

func NewCoordinate(x int, y int) Coordinate {
	return Coordinate{
		x: x,
		y: y,
	}
}

func (c Coordinate) GetX() int {
	return c.x
}

func (c Coordinate) GetY() int {
	return c.y
}
