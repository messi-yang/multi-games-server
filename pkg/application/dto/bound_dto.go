package dto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
)

type BoundDto struct {
	From LocationDto `json:"from"`
	To   LocationDto `json:"to"`
}

func NewBoundDto(bound commonmodel.BoundVo) BoundDto {
	return BoundDto{
		From: NewLocationDto(bound.GetFrom()),
		To:   NewLocationDto(bound.GetTo()),
	}
}

func (dto BoundDto) ToValueObject() (commonmodel.BoundVo, error) {
	from := commonmodel.NewLocationVo(dto.From.X, dto.From.Z)

	to := commonmodel.NewLocationVo(dto.To.X, dto.To.Z)

	bound, err := commonmodel.NewBoundVo(
		from,
		to,
	)
	if err != nil {
		return commonmodel.BoundVo{}, err
	}

	return bound, nil
}
