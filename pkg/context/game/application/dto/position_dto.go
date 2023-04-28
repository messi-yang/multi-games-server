package dto

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"

type PositionDto struct {
	X int `json:"x"`
	Z int `json:"z"`
}

func NewPositionDto(position commonmodel.Position) PositionDto {
	return PositionDto{
		X: position.GetX(),
		Z: position.GetZ(),
	}
}
