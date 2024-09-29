package dto

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

type DimensionDto struct {
	Width int `json:"width"`
	Depth int `json:"depth"`
}

func NewDimensionDto(dimension worldcommonmodel.Dimension) DimensionDto {
	return DimensionDto{
		Width: dimension.GetWidth(),
		Depth: dimension.GetDepth(),
	}
}
