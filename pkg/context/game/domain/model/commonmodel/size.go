package commonmodel

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/domainmodel"
)

type ErrInvalidSize struct {
	width  int
	height int
}

func (e *ErrInvalidSize) Error() string {
	return fmt.Sprintf("width (%d) and height(%d) of size must be greater than 0", e.width, e.height)
}

type Size struct {
	width  int
	height int
}

// Interface Implementation Check
var _ domainmodel.ValueObject[Size] = (*Size)(nil)

func NewSize(width int, height int) (Size, error) {
	if width < 1 || height < 1 {
		return Size{}, &ErrInvalidSize{width: width, height: height}
	}

	return Size{
		width:  width,
		height: height,
	}, nil
}

func (size Size) IsEqual(otherSize Size) bool {
	return size.width == otherSize.width && size.height == otherSize.height
}

func (size Size) GetWidth() int {
	return size.width
}

func (size Size) GetHeight() int {
	return size.height
}
