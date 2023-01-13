package commonmodel

import (
	"fmt"

	"github.com/samber/lo"
)

type ErrInvalidBound struct {
	from Location
	to   Location
}

func (e *ErrInvalidBound) Error() string {
	return fmt.Sprintf("from location (%+v) cannot exceed to location (%+v)", e.from, e.to)
}

type Bound struct {
	from Location
	to   Location
}

func NewBound(from Location, to Location) (Bound, error) {
	if from.x > to.x || from.y > to.y {
		return Bound{}, &ErrInvalidBound{from: from, to: to}
	}

	return Bound{
		from: from,
		to:   to,
	}, nil
}

func (a Bound) GetFrom() Location {
	return a.from
}

func (a Bound) GetTo() Location {
	return a.to
}

func (a Bound) GetWidth() int {
	return a.to.x - a.from.x + 1
}

func (a Bound) GetHeight() int {
	return a.to.y - a.from.y + 1
}

func (a Bound) IncludesLocation(location Location) bool {
	return location.x >= a.from.x && location.x <= a.to.x && location.y >= a.from.y && location.y <= a.to.y
}

func (a Bound) IncludesAnyLocations(locations []Location) bool {
	return lo.ContainsBy(locations, func(location Location) bool {
		return a.IncludesLocation(location)
	})
}
