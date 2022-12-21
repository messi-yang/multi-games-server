package commonmodel

import "fmt"

type ErrInvalidCoordinate struct {
	x int
	y int
}

func (e *ErrInvalidCoordinate) Error() string {
	return fmt.Sprintf("coordinate (%d, %d) cannot have negative values", e.x, e.y)
}

type Coordinate struct {
	x int
	y int
}

func NewCoordinate(x int, y int) (Coordinate, error) {
	if x < 0 || y < 0 {
		return Coordinate{}, &ErrInvalidCoordinate{x: x, y: y}
	}
	return Coordinate{
		x: x,
		y: y,
	}, nil
}

func (c Coordinate) GetX() int {
	return c.x
}

func (c Coordinate) GetY() int {
	return c.y
}
