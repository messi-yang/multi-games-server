package commonmodel

import (
	"fmt"

	"github.com/samber/lo"
)

type ErrInvalidRangeVo struct {
	from Location
	to   Location
}

func (e *ErrInvalidRangeVo) Error() string {
	return fmt.Sprintf("from location (%+v) cannot exceed to location (%+v)", e.from, e.to)
}

type RangeVo struct {
	from Location
	to   Location
}

func NewRangeVo(from Location, to Location) (RangeVo, error) {
	if from.x > to.x || from.y > to.y {
		return RangeVo{}, &ErrInvalidRangeVo{from: from, to: to}
	}

	return RangeVo{
		from: from,
		to:   to,
	}, nil
}

func (a RangeVo) GetFrom() Location {
	return a.from
}

func (a RangeVo) GetTo() Location {
	return a.to
}

func (a RangeVo) GetWidth() int {
	return a.to.x - a.from.x + 1
}

func (a RangeVo) GetHeight() int {
	return a.to.y - a.from.y + 1
}

func (a RangeVo) IncludesLocation(location Location) bool {
	return location.x >= a.from.x && location.x <= a.to.x && location.y >= a.from.y && location.y <= a.to.y
}

func (a RangeVo) IncludesAnyLocations(locations []Location) bool {
	return lo.ContainsBy(locations, func(location Location) bool {
		return a.IncludesLocation(location)
	})
}
