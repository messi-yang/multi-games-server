package uidto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamecommonmodel"
)

type DimensionUiDto struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func NewDimensionUiDto(dimension gamecommonmodel.Dimension) DimensionUiDto {
	return DimensionUiDto{
		Width:  dimension.GetWidth(),
		Height: dimension.GetHeight(),
	}
}

func (dto DimensionUiDto) ToValueObject() (gamecommonmodel.Dimension, error) {
	return gamecommonmodel.NewDimension(dto.Width, dto.Height)
}
