package jsondto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
)

type BoundVoDto struct {
	From PositionVoDto `json:"from"`
	To   PositionVoDto `json:"to"`
}

func NewBoundVoDto(bound commonmodel.BoundVo) BoundVoDto {
	return BoundVoDto{
		From: NewPositionVoDto(bound.GetFrom()),
		To:   NewPositionVoDto(bound.GetTo()),
	}
}

func (dto BoundVoDto) ToValueObject() (bound commonmodel.BoundVo, err error) {
	return commonmodel.NewBoundVo(
		commonmodel.NewPositionVo(dto.From.X, dto.From.Z),
		commonmodel.NewPositionVo(dto.To.X, dto.To.Z),
	)
}
