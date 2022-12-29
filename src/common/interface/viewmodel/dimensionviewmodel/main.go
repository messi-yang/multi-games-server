package dimensionviewmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"

type DimensionViewModel struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func New(dimension commonmodel.Dimension) DimensionViewModel {
	return DimensionViewModel{
		Width:  dimension.GetWidth(),
		Height: dimension.GetHeight(),
	}
}

func (dto DimensionViewModel) ToValueObject() (commonmodel.Dimension, error) {
	return commonmodel.NewDimension(dto.Width, dto.Height)
}
