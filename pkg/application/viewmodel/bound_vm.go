package viewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
)

type BoundVm struct {
	From LocationVm `json:"from"`
	To   LocationVm `json:"to"`
}

func NewBoundVm(bound commonmodel.BoundVo) BoundVm {
	return BoundVm{
		From: NewLocationVm(bound.GetFrom()),
		To:   NewLocationVm(bound.GetTo()),
	}
}

func (dto BoundVm) ToValueObject() (commonmodel.BoundVo, error) {
	from := commonmodel.NewLocationVo(dto.From.X, dto.From.Y)

	to := commonmodel.NewLocationVo(dto.To.X, dto.To.Y)

	bound, err := commonmodel.NewBoundVo(
		from,
		to,
	)
	if err != nil {
		return commonmodel.BoundVo{}, err
	}

	return bound, nil
}
