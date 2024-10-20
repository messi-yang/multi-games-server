package dto

import "github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"

type PositionDto struct {
	X int `json:"x"`
	Z int `json:"z"`
}

func NewPositionDto(position worldcommonmodel.Position) PositionDto {
	return PositionDto{
		X: position.GetX(),
		Z: position.GetZ(),
	}
}

func (dto PositionDto) ToValueObject() (bound worldcommonmodel.Position) {
	return worldcommonmodel.NewPosition(
		dto.X,
		dto.Z,
	)
}
