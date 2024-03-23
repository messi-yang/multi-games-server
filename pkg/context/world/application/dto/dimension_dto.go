package dto

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

type DimensionDto struct {
	Width int8 `json:"width"`
	Depth int8 `json:"depth"`
}

func NewDimensionDto(dimension worldcommonmodel.Dimension) DimensionDto {
	return DimensionDto{
		Width: dimension.GetWidth(),
		Depth: dimension.GetDepth(),
	}
}
