package dto

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/commonmodel"
)

type BoundDto struct {
	From PositionDto `json:"from"`
	To   PositionDto `json:"to"`
}

func NewBoundDto(bound commonmodel.Bound) BoundDto {
	return BoundDto{
		From: NewPositionDto(bound.GetFrom()),
		To:   NewPositionDto(bound.GetTo()),
	}
}

func (dto BoundDto) ToValueObject() (bound commonmodel.Bound, err error) {
	return commonmodel.NewBound(
		commonmodel.NewPosition(dto.From.X, dto.From.Z),
		commonmodel.NewPosition(dto.To.X, dto.To.Z),
	)
}
