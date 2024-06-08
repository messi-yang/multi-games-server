package dto

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

type BoundDto struct {
	From PositionDto `json:"from"`
	To   PositionDto `json:"to"`
}

func NewBoundDto(bound worldcommonmodel.Bound) BoundDto {
	return BoundDto{
		From: NewPositionDto(bound.GetFrom()),
		To:   NewPositionDto(bound.GetTo()),
	}
}

func (dto BoundDto) ToValueObject() (bound worldcommonmodel.Bound, err error) {
	return worldcommonmodel.NewBound(
		worldcommonmodel.NewPosition(dto.From.X, dto.From.Z),
		worldcommonmodel.NewPosition(dto.To.X, dto.To.Z),
	)
}
