package commonmodel

import (
	"fmt"

	"github.com/samber/lo"
)

type ErrInvalidMapRange struct {
	from Location
	to   Location
}

func (e *ErrInvalidMapRange) Error() string {
	return fmt.Sprintf("from location (%+v) cannot exceed to location (%+v)", e.from, e.to)
}

type MapRange struct {
	from Location
	to   Location
}

func NewMapRange(from Location, to Location) (MapRange, error) {
	if from.x > to.x || from.y > to.y {
		return MapRange{}, &ErrInvalidMapRange{from: from, to: to}
	}

	return MapRange{
		from: from,
		to:   to,
	}, nil
}

func (a MapRange) GetFrom() Location {
	return a.from
}

func (a MapRange) GetTo() Location {
	return a.to
}

func (a MapRange) GetWidth() int {
	return a.to.x - a.from.x + 1
}

func (a MapRange) GetHeight() int {
	return a.to.y - a.from.y + 1
}

func (a MapRange) IncludesLocation(location Location) bool {
	return location.x >= a.from.x && location.x <= a.to.x && location.y >= a.from.y && location.y <= a.to.y
}

func (a MapRange) IncludesAnyLocations(locations []Location) bool {
	return lo.ContainsBy(locations, func(location Location) bool {
		return a.IncludesLocation(location)
	})
}
