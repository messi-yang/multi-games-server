package viewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/gamemodel"
)

type BoundVm struct {
	From LocationVm `json:"from"`
	To   LocationVm `json:"to"`
}

func NewBoundVm(bound gamemodel.BoundVo) BoundVm {
	return BoundVm{
		From: NewLocationVm(bound.GetFrom()),
		To:   NewLocationVm(bound.GetTo()),
	}
}

func (dto BoundVm) ToValueObject() (gamemodel.BoundVo, error) {
	from := commonmodel.NewLocationVo(dto.From.X, dto.From.Y)

	to := commonmodel.NewLocationVo(dto.To.X, dto.To.Y)

	bound, err := gamemodel.NewBoundVo(
		from,
		to,
	)
	if err != nil {
		return gamemodel.BoundVo{}, err
	}

	return bound, nil
}
