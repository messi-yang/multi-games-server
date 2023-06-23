package commonmodel

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain"
)

type ErrInvalidBound struct {
	from Position
	to   Position
}

func (e *ErrInvalidBound) Error() string {
	return fmt.Sprintf("from position (%+v) cannot exceed to position (%+v)", e.from, e.to)
}

type Bound struct {
	from Position
	to   Position
}

// Interface Implementation Check
var _ domain.ValueObject[Bound] = (*Bound)(nil)

func NewBound(from Position, to Position) (Bound, error) {
	if from.GetX() > to.GetX() || from.GetZ() > to.GetZ() {
		return Bound{}, &ErrInvalidBound{from: from, to: to}
	}

	return Bound{
		from: from,
		to:   to,
	}, nil
}

func (bound Bound) IsEqual(otherBound Bound) bool {
	return bound.from.IsEqual(otherBound.from) && bound.to.IsEqual(otherBound.to)
}

func (bound Bound) GetFrom() Position {
	return bound.from
}

func (bound Bound) GetTo() Position {
	return bound.to
}

func (bound Bound) GetWidth() int {
	return bound.to.GetX() - bound.from.GetX() + 1
}

func (bound Bound) GetHeight() int {
	return bound.to.GetZ() - bound.from.GetZ() + 1
}

func (bound Bound) CoversPosition(position Position) bool {
	return position.GetX() >= bound.from.GetX() && position.GetX() <= bound.to.GetX() && position.GetZ() >= bound.from.GetZ() && position.GetZ() <= bound.to.GetZ()
}
