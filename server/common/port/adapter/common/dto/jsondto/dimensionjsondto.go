package jsondto

import (
	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
)

type DimensionJsonDto struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func NewDimensionJsonDto(dimension gamecommonmodel.Dimension) DimensionJsonDto {
	return DimensionJsonDto{
		Width:  dimension.GetWidth(),
		Height: dimension.GetHeight(),
	}
}

func (dto DimensionJsonDto) ToValueObject() (gamecommonmodel.Dimension, error) {
	return gamecommonmodel.NewDimension(dto.Width, dto.Height)
}
