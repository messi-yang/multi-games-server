package gamecommonmodel

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
)

type ErrInvalidDimension struct {
	width int
	depth int
}

func (e *ErrInvalidDimension) Error() string {
	return fmt.Sprintf("width (%+v) or depth (%+v) cannot be less than 1", e.width, e.depth)
}

type Dimension struct {
	width int
	depth int
}

// Interface Implementation Check
var _ domain.ValueObject[Dimension] = (*Dimension)(nil)

func NewDimension(width int, depth int) (Dimension, error) {
	if width < 1 || depth < 1 {
		return Dimension{}, &ErrInvalidDimension{width: width, depth: depth}
	}

	return Dimension{
		width: width,
		depth: depth,
	}, nil
}

func (dimension Dimension) IsEqual(otherDimension Dimension) bool {
	return dimension.width == otherDimension.width && dimension.depth == otherDimension.depth
}

func (dimension Dimension) GetWidth() int {
	return dimension.width
}

func (dimension Dimension) GetDepth() int {
	return dimension.depth
}

func (dimension Dimension) IsSymmetric() bool {
	return dimension.width == dimension.depth
}
