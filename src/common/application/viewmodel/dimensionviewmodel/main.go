package dimensionviewmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"

type ViewModel struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func New(dimension commonmodel.Dimension) ViewModel {
	return ViewModel{
		Width:  dimension.GetWidth(),
		Height: dimension.GetHeight(),
	}
}

func (dto ViewModel) ToValueObject() (commonmodel.Dimension, error) {
	return commonmodel.NewDimension(dto.Width, dto.Height)
}
