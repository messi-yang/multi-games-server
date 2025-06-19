package dto

import "github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/gamecommonmodel"

type PositionDto struct {
	X int `json:"x"`
	Z int `json:"z"`
}

func NewPositionDto(position gamecommonmodel.Position) PositionDto {
	return PositionDto{
		X: position.GetX(),
		Z: position.GetZ(),
	}
}

func (dto PositionDto) ToValueObject() (bound gamecommonmodel.Position) {
	return gamecommonmodel.NewPosition(
		dto.X,
		dto.Z,
	)
}
