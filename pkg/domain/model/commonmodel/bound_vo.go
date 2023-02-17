package commonmodel

import (
	"fmt"

	"github.com/samber/lo"
)

type ErrInvalidBoundVo struct {
	from LocationVo
	to   LocationVo
}

func (e *ErrInvalidBoundVo) Error() string {
	return fmt.Sprintf("from location (%+v) cannot exceed to location (%+v)", e.from, e.to)
}

type BoundVo struct {
	from LocationVo
	to   LocationVo
}

func NewBoundVo(from LocationVo, to LocationVo) (BoundVo, error) {
	if from.GetX() > to.GetX() || from.GetZ() > to.GetZ() {
		return BoundVo{}, &ErrInvalidBoundVo{from: from, to: to}
	}

	return BoundVo{
		from: from,
		to:   to,
	}, nil
}

func (bound BoundVo) GetFrom() LocationVo {
	return bound.from
}

func (bound BoundVo) GetTo() LocationVo {
	return bound.to
}

func (bound BoundVo) GetWidth() int {
	return bound.to.GetX() - bound.from.GetX() + 1
}

func (bound BoundVo) GetHeight() int {
	return bound.to.GetZ() - bound.from.GetZ() + 1
}

func (bound BoundVo) CoversLocation(location LocationVo) bool {
	return location.GetX() >= bound.from.GetX() && location.GetX() <= bound.to.GetX() && location.GetZ() >= bound.from.GetZ() && location.GetZ() <= bound.to.GetZ()
}

func (bound BoundVo) CoverAnyLocations(locations []LocationVo) bool {
	return lo.ContainsBy(locations, func(location LocationVo) bool {
		return bound.CoversLocation(location)
	})
}
