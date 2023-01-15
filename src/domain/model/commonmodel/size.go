package commonmodel

import (
	"fmt"
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

func NewSize(width int, height int) (Size, error) {
	if width < 1 || height < 1 {
		return Size{}, &ErrInvalidSize{width: width, height: height}
	}

	return Size{
		width:  width,
		height: height,
	}, nil
}

func (ms Size) GetWidth() int {
	return ms.width
}

func (ms Size) GetHeight() int {
	return ms.height
}

func (ms Size) CoverLocation(location Location) bool {
	includesAll := true
	if location.x >= ms.width || location.y >= ms.height {
		includesAll = false
	}

	return includesAll
}
