package dto

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"

type PositionVoDto struct {
	X int `json:"x"`
	Z int `json:"z"`
}

func NewPositionVoDto(position commonmodel.PositionVo) PositionVoDto {
	return PositionVoDto{
		X: position.GetX(),
		Z: position.GetZ(),
	}
}
