package viewmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"

type DimensionVm struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func NewDimensionVm(dimension commonmodel.Dimension) DimensionVm {
	return DimensionVm{
		Width:  dimension.GetWidth(),
		Height: dimension.GetHeight(),
	}
}

func (dto DimensionVm) ToValueObject() (commonmodel.Dimension, error) {
	return commonmodel.NewDimension(dto.Width, dto.Height)
}
