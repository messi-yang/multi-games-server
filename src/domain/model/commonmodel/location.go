package commonmodel

import "fmt"

type ErrInvalidLocation struct {
	x int
	y int
}

func (e *ErrInvalidLocation) Error() string {
	return fmt.Sprintf("location (%d, %d) cannot have negative values", e.x, e.y)
}

type Location struct {
	x int
	y int
}

func NewLocation(x int, y int) (Location, error) {
	if x < 0 || y < 0 {
		return Location{}, &ErrInvalidLocation{x: x, y: y}
	}
	return Location{
		x: x,
		y: y,
	}, nil
}

func (c Location) GetX() int {
	return c.x
}

func (c Location) GetY() int {
	return c.y
}
