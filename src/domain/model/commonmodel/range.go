package commonmodel

import (
	"fmt"

	"github.com/samber/lo"
)

type ErrInvalidRange struct {
	from Location
	to   Location
}

func (e *ErrInvalidRange) Error() string {
	return fmt.Sprintf("from location (%+v) cannot exceed to location (%+v)", e.from, e.to)
}

type Range struct {
	from Location
	to   Location
}

func NewRange(from Location, to Location) (Range, error) {
	if from.x > to.x || from.y > to.y {
		return Range{}, &ErrInvalidRange{from: from, to: to}
	}

	return Range{
		from: from,
		to:   to,
	}, nil
}

func (a Range) GetFrom() Location {
	return a.from
}

func (a Range) GetTo() Location {
	return a.to
}

func (a Range) GetWidth() int {
	return a.to.x - a.from.x + 1
}

func (a Range) GetHeight() int {
	return a.to.y - a.from.y + 1
}

func (a Range) IncludesLocation(location Location) bool {
	return location.x >= a.from.x && location.x <= a.to.x && location.y >= a.from.y && location.y <= a.to.y
}

func (a Range) IncludesAnyLocations(locations []Location) bool {
	return lo.ContainsBy(locations, func(location Location) bool {
		return a.IncludesLocation(location)
	})
}
