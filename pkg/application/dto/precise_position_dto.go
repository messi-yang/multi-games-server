package dto

import "github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"

type PrecisePositionDto struct {
	X float32 `json:"x"`
	Z float32 `json:"z"`
}

func NewPrecisePositionDto(precisePosition worldcommonmodel.PrecisePosition) PrecisePositionDto {
	return PrecisePositionDto{
		X: precisePosition.GetX(),
		Z: precisePosition.GetZ(),
	}
}
