package dto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamecommonmodel"
)

type DimensionDto struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func NewDimensionDto(dimension gamecommonmodel.Dimension) DimensionDto {
	return DimensionDto{
		Width:  dimension.GetWidth(),
		Height: dimension.GetHeight(),
	}
}

func (dto DimensionDto) ToValueObject() (gamecommonmodel.Dimension, error) {
	return gamecommonmodel.NewDimension(dto.Width, dto.Height)
}
