package commonmodel

import (
	"fmt"

	"github.com/samber/lo"
)

type ErrInvalidBoundVo struct {
	from PositionVo
	to   PositionVo
}

func (e *ErrInvalidBoundVo) Error() string {
	return fmt.Sprintf("from position (%+v) cannot exceed to position (%+v)", e.from, e.to)
}

type BoundVo struct {
	from PositionVo
	to   PositionVo
}

func NewBoundVo(from PositionVo, to PositionVo) (BoundVo, error) {
	if from.GetX() > to.GetX() || from.GetZ() > to.GetZ() {
		return BoundVo{}, &ErrInvalidBoundVo{from: from, to: to}
	}

	return BoundVo{
		from: from,
		to:   to,
	}, nil
}

func (bound BoundVo) GetFrom() PositionVo {
	return bound.from
}

func (bound BoundVo) GetTo() PositionVo {
	return bound.to
}

func (bound BoundVo) GetWidth() int {
	return bound.to.GetX() - bound.from.GetX() + 1
}

func (bound BoundVo) GetHeight() int {
	return bound.to.GetZ() - bound.from.GetZ() + 1
}

func (bound BoundVo) GetCenterPos() PositionVo {
	return NewPositionVo((bound.from.x+bound.to.x)/2, (bound.from.z+bound.to.z)/2)
}

func (bound BoundVo) CoversPosition(position PositionVo) bool {
	return position.GetX() >= bound.from.GetX() && position.GetX() <= bound.to.GetX() && position.GetZ() >= bound.from.GetZ() && position.GetZ() <= bound.to.GetZ()
}

func (bound BoundVo) CoverAnyPositions(positions []PositionVo) bool {
	return lo.ContainsBy(positions, func(position PositionVo) bool {
		return bound.CoversPosition(position)
	})
}
