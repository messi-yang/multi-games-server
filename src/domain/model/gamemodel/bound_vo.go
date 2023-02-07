package gamemodel

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/samber/lo"
)

type ErrInvalidBoundVo struct {
	from commonmodel.LocationVo
	to   commonmodel.LocationVo
}

func (e *ErrInvalidBoundVo) Error() string {
	return fmt.Sprintf("from location (%+v) cannot exceed to location (%+v)", e.from, e.to)
}

type BoundVo struct {
	from commonmodel.LocationVo
	to   commonmodel.LocationVo
}

func NewBoundVo(from commonmodel.LocationVo, to commonmodel.LocationVo) (BoundVo, error) {
	if from.GetX() > to.GetX() || from.GetY() > to.GetY() {
		return BoundVo{}, &ErrInvalidBoundVo{from: from, to: to}
	}

	return BoundVo{
		from: from,
		to:   to,
	}, nil
}

func (bound BoundVo) GetFrom() commonmodel.LocationVo {
	return bound.from
}

func (bound BoundVo) GetTo() commonmodel.LocationVo {
	return bound.to
}

func (bound BoundVo) GetWidth() int {
	return bound.to.GetX() - bound.from.GetX() + 1
}

func (bound BoundVo) GetHeight() int {
	return bound.to.GetY() - bound.from.GetY() + 1
}

func (bound BoundVo) CoversLocation(location commonmodel.LocationVo) bool {
	return location.GetX() >= bound.from.GetX() && location.GetX() <= bound.to.GetX() && location.GetY() >= bound.from.GetY() && location.GetY() <= bound.to.GetY()
}

func (bound BoundVo) CoverAnyLocations(locations []commonmodel.LocationVo) bool {
	return lo.ContainsBy(locations, func(location commonmodel.LocationVo) bool {
		return bound.CoversLocation(location)
	})
}
