package worldcommonmodel

import "github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"

type PrecisePosition struct {
	x float32
	z float32
}

// Interface Implementation Check
var _ domain.ValueObject[PrecisePosition] = (*PrecisePosition)(nil)

func NewPrecisePosition(x float32, z float32) PrecisePosition {
	return PrecisePosition{
		x: x,
		z: z,
	}
}

func (position PrecisePosition) IsEqual(otherPosition PrecisePosition) bool {
	return position.x == otherPosition.x && position.z == otherPosition.z
}

func (position PrecisePosition) GetX() float32 {
	return position.x
}

func (position PrecisePosition) GetZ() float32 {
	return position.z
}

func (position PrecisePosition) Shift(x float32, z float32) PrecisePosition {
	return NewPrecisePosition(position.x+x, position.z+z)
}
