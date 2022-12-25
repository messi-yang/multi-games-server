package jsondto

import "github.com/dum-dum-genius/game-of-liberty-computer/domain/domainmodel/commonmodel"

type DimensionJsonDto struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func NewDimensionJsonDto(dimension commonmodel.Dimension) DimensionJsonDto {
	return DimensionJsonDto{
		Width:  dimension.GetWidth(),
		Height: dimension.GetHeight(),
	}
}

func (dto DimensionJsonDto) ToValueObject() (commonmodel.Dimension, error) {
	return commonmodel.NewDimension(dto.Width, dto.Height)
}
