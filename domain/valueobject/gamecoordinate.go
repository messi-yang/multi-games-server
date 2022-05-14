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
