package presenterdto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamecommonmodel"
)

type DimensionPresenterDto struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func NewDimensionPresenterDto(dimension gamecommonmodel.Dimension) DimensionPresenterDto {
	return DimensionPresenterDto{
		Width:  dimension.GetWidth(),
		Height: dimension.GetHeight(),
	}
}

func (dto DimensionPresenterDto) ToValueObject() (gamecommonmodel.Dimension, error) {
	return gamecommonmodel.NewDimension(dto.Width, dto.Height)
}
