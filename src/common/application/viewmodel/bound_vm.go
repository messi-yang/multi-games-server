package viewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
)

type BoundVm struct {
	From LocationVm `json:"from"`
	To   LocationVm `json:"to"`
}

func NewBoundVm(bound commonmodel.Bound) BoundVm {
	return BoundVm{
		From: NewLocationVm(bound.GetFrom()),
		To:   NewLocationVm(bound.GetTo()),
	}
}

func (dto BoundVm) ToValueObject() (commonmodel.Bound, error) {
	from, err := commonmodel.NewLocation(dto.From.X, dto.From.Y)
	if err != nil {
		return commonmodel.Bound{}, err
	}

	to, err := commonmodel.NewLocation(dto.To.X, dto.To.Y)
	if err != nil {
		return commonmodel.Bound{}, err
	}

	bound, err := commonmodel.NewBound(
		from,
		to,
	)
	if err != nil {
		return commonmodel.Bound{}, err
	}

	return bound, nil
}
