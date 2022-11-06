package dto

import "github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game/valueobject"

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
