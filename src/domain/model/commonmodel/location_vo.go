package commonmodel

import "fmt"

type ErrInvalidLocationVo struct {
	x int
	y int
}

func (e *ErrInvalidLocationVo) Error() string {
	return fmt.Sprintf("location (%d, %d) cannot have negative values", e.x, e.y)
}

type LocationVo struct {
	x int
	y int
}

func NewLocationVo(x int, y int) (LocationVo, error) {
	if x < 0 || y < 0 {
		return LocationVo{}, &ErrInvalidLocationVo{x: x, y: y}
	}
	return LocationVo{
		x: x,
		y: y,
	}, nil
}

func (c LocationVo) GetX() int {
	return c.x
}

func (c LocationVo) GetY() int {
	return c.y
}
