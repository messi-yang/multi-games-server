package commonmodel

import (
	"fmt"

	"github.com/samber/lo"
)

type ErrInvalidExtent struct {
	from Location
	to   Location
}

func (e *ErrInvalidExtent) Error() string {
	return fmt.Sprintf("from location (%+v) cannot exceed to location (%+v)", e.from, e.to)
}

type Extent struct {
	from Location
	to   Location
}

func NewExtent(from Location, to Location) (Extent, error) {
	if from.x > to.x || from.y > to.y {
		return Extent{}, &ErrInvalidExtent{from: from, to: to}
	}

	return Extent{
		from: from,
		to:   to,
	}, nil
}

func (a Extent) GetFrom() Location {
	return a.from
}

func (a Extent) GetTo() Location {
	return a.to
}

func (a Extent) GetWidth() int {
	return a.to.x - a.from.x + 1
}

func (a Extent) GetHeight() int {
	return a.to.y - a.from.y + 1
}

func (a Extent) IncludesLocation(location Location) bool {
	return location.x >= a.from.x && location.x <= a.to.x && location.y >= a.from.y && location.y <= a.to.y
}

func (a Extent) IncludesAnyLocations(locations []Location) bool {
	return lo.ContainsBy(locations, func(location Location) bool {
		return a.IncludesLocation(location)
	})
}
