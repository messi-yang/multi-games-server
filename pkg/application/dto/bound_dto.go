package dto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
)

type BoundDto struct {
	From PositionDto `json:"from"`
	To   PositionDto `json:"to"`
}

func NewBoundDto(bound commonmodel.BoundVo) BoundDto {
	return BoundDto{
		From: NewPositionDto(bound.GetFrom()),
		To:   NewPositionDto(bound.GetTo()),
	}
}

func (dto BoundDto) ToValueObject() (commonmodel.BoundVo, error) {
	from := commonmodel.NewPositionVo(dto.From.X, dto.From.Z)

	to := commonmodel.NewPositionVo(dto.To.X, dto.To.Z)

	bound, err := commonmodel.NewBoundVo(
		from,
		to,
	)
	if err != nil {
		return commonmodel.BoundVo{}, err
	}

	return bound, nil
}
