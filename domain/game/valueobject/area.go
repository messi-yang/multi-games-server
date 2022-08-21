package valueobject

import "fmt"

type ErrInvalidArea struct {
	from Coordinate
	to   Coordinate
}

func (e *ErrInvalidArea) Error() string {
	return fmt.Sprintf("from coordinate (%+v) cannot exceed to coordinate (%+v)", e.from, e.to)
}

type Area struct {
	from Coordinate
	to   Coordinate
}

func NewArea(from Coordinate, to Coordinate) (Area, error) {
	if from.x > to.x || from.y > to.y {
		return Area{}, &ErrInvalidArea{from: from, to: to}
	}

	return Area{
		from: from,
		to:   to,
	}, nil
}

func (a Area) GetFrom() Coordinate {
	return a.from
}

func (a Area) GetTo() Coordinate {
	return a.to
}

func (a Area) IncludesCoordinate(coordinate Coordinate) bool {
	return coordinate.x >= a.from.x && coordinate.x <= a.to.x && coordinate.y >= a.from.y && coordinate.y <= a.to.y
}
