package commonmodel

import (
	"errors"

	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain"
)

var (
	ErrInvalidSize = errors.New("width or height is invalid")
)

type Size struct {
	width  int
	height int
}

// Interface Implementation Check
var _ domain.ValueObject[Size] = (*Size)(nil)

func NewSize(width int, height int) (Size, error) {
	if width < 1 || height < 1 {
		return Size{}, ErrInvalidSize
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
