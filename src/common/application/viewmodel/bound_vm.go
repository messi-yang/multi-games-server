package viewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/livegamemodel"
)

type BoundVm struct {
	From LocationVm `json:"from"`
	To   LocationVm `json:"to"`
}

func NewBoundVm(bound livegamemodel.Bound) BoundVm {
	return BoundVm{
		From: NewLocationVm(bound.GetFrom()),
		To:   NewLocationVm(bound.GetTo()),
	}
}

func (dto BoundVm) ToValueObject() (livegamemodel.Bound, error) {
	from, err := commonmodel.NewLocation(dto.From.X, dto.From.Y)
	if err != nil {
		return livegamemodel.Bound{}, err
	}

	to, err := commonmodel.NewLocation(dto.To.X, dto.To.Y)
	if err != nil {
		return livegamemodel.Bound{}, err
	}

	bound, err := livegamemodel.NewBound(
		from,
		to,
	)
	if err != nil {
		return livegamemodel.Bound{}, err
	}

	return bound, nil
}
