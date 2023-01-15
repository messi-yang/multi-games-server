package livegamemodel

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/samber/lo"
)

type ErrInvalidBound struct {
	from commonmodel.Location
	to   commonmodel.Location
}

func (e *ErrInvalidBound) Error() string {
	return fmt.Sprintf("from location (%+v) cannot exceed to location (%+v)", e.from, e.to)
}

type Bound struct {
	from commonmodel.Location
	to   commonmodel.Location
}

func NewBound(from commonmodel.Location, to commonmodel.Location) (Bound, error) {
	if from.GetX() > to.GetX() || from.GetY() > to.GetY() {
		return Bound{}, &ErrInvalidBound{from: from, to: to}
	}

	return Bound{
		from: from,
		to:   to,
	}, nil
}

func (a Bound) GetFrom() commonmodel.Location {
	return a.from
}

func (a Bound) GetTo() commonmodel.Location {
	return a.to
}

func (a Bound) GetWidth() int {
	return a.to.GetX() - a.from.GetX() + 1
}

func (a Bound) GetHeight() int {
	return a.to.GetY() - a.from.GetY() + 1
}

func (a Bound) CoverLocation(location commonmodel.Location) bool {
	return location.GetX() >= a.from.GetX() && location.GetX() <= a.to.GetX() && location.GetY() >= a.from.GetY() && location.GetY() <= a.to.GetY()
}

func (a Bound) CoverAnyLocations(locations []commonmodel.Location) bool {
	return lo.ContainsBy(locations, func(location commonmodel.Location) bool {
		return a.CoverLocation(location)
	})
}
