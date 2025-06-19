package dto

import "github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/gamecommonmodel"

type PrecisePositionDto struct {
	X float32 `json:"x"`
	Z float32 `json:"z"`
}

func NewPrecisePositionDto(precisePosition gamecommonmodel.PrecisePosition) PrecisePositionDto {
	return PrecisePositionDto{
		X: precisePosition.GetX(),
		Z: precisePosition.GetZ(),
	}
}
