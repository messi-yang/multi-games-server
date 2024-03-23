package worldcommonmodel

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
)

type ErrInvalidDimension struct {
	width int8
	depth int8
}

func (e *ErrInvalidDimension) Error() string {
	return fmt.Sprintf("width (%+v) or depth (%+v) cannot be less than 1", e.width, e.depth)
}

type Dimension struct {
	width int8
	depth int8
}

// Interface Implementation Check
var _ domain.ValueObject[Dimension] = (*Dimension)(nil)

func NewDimension(width int8, depth int8) (Dimension, error) {
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

func (dimension Dimension) GetWidth() int8 {
	return dimension.width
}

func (dimension Dimension) GetDepth() int8 {
	return dimension.depth
}
