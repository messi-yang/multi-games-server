package dto

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"

type PositionDto struct {
	X int `json:"x"`
	Z int `json:"z"`
}

func NewPositionDto(position commonmodel.PositionVo) PositionDto {
	return PositionDto{
		X: position.GetX(),
		Z: position.GetZ(),
	}
}
