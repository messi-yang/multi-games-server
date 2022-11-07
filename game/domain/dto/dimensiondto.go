package dto

import "github.com/dum-dum-genius/game-of-liberty-computer/game/domain/valueobject"

type DimensionDto struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func NewDimensionDto(dimension valueobject.Dimension) DimensionDto {
	return DimensionDto{
		Width:  dimension.GetWidth(),
		Height: dimension.GetHeight(),
	}
}

func (dto DimensionDto) ToValueObject() (valueobject.Dimension, error) {
	return valueobject.NewDimension(dto.Width, dto.Height)
}
