package viewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
)

type BoundVm struct {
	From LocationVm `json:"from"`
	To   LocationVm `json:"to"`
}

func NewBoundVm(bound_ commonmodel.Bound) BoundVm {
	return BoundVm{
		From: NewLocationVm(bound_.GetFrom()),
		To:   NewLocationVm(bound_.GetTo()),
	}
}

func (dto BoundVm) ToValueObject() (commonmodel.Bound, error) {
	fromLocationVm, err := dto.From.ToValueObject()
	if err != nil {
		return commonmodel.Bound{}, err
	}

	toLocationVm, err := dto.To.ToValueObject()
	if err != nil {
		return commonmodel.Bound{}, err
	}

	bound_, err := commonmodel.NewBound(
		fromLocationVm,
		toLocationVm,
	)
	if err != nil {
		return commonmodel.Bound{}, err
	}

	return bound_, nil
}
