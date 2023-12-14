package worldcommonmodel

import (
	"math"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
)

type PrecisePosition struct {
	x float32
	z float32
}

// Interface Implementation Check
var _ domain.ValueObject[PrecisePosition] = (*PrecisePosition)(nil)

func NewPrecisePosition(x float32, z float32) PrecisePosition {
	return PrecisePosition{
		x: float32(math.Round(float64(x*100)) / 100),
		z: float32(math.Round(float64(x*100)) / 100),
	}
}

func (precisePosition PrecisePosition) IsEqual(otherPosition PrecisePosition) bool {
	return precisePosition.x == otherPosition.x && precisePosition.z == otherPosition.z
}

func (precisePosition PrecisePosition) GetX() float32 {
	return precisePosition.x
}

func (precisePosition PrecisePosition) GetZ() float32 {
	return precisePosition.z
}

func (precisePosition PrecisePosition) Shift(x float32, z float32) PrecisePosition {
	return NewPrecisePosition(precisePosition.x+x, precisePosition.z+z)
}
