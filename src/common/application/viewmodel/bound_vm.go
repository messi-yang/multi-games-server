package viewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/livegamemodel"
)

type BoundVm struct {
	From LocationVm `json:"from"`
	To   LocationVm `json:"to"`
}

func NewBoundVm(bound livegamemodel.BoundVo) BoundVm {
	return BoundVm{
		From: NewLocationVm(bound.GetFrom()),
		To:   NewLocationVm(bound.GetTo()),
	}
}

func (dto BoundVm) ToValueObject() (livegamemodel.BoundVo, error) {
	from, err := commonmodel.NewLocationVo(dto.From.X, dto.From.Y)
	if err != nil {
		return livegamemodel.BoundVo{}, err
	}

	to, err := commonmodel.NewLocationVo(dto.To.X, dto.To.Y)
	if err != nil {
		return livegamemodel.BoundVo{}, err
	}

	bound, err := livegamemodel.NewBoundVo(
		from,
		to,
	)
	if err != nil {
		return livegamemodel.BoundVo{}, err
	}

	return bound, nil
}
