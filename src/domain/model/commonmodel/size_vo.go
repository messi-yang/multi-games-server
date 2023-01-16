package commonmodel

import (
	"fmt"
)

type ErrInvalidSizeVo struct {
	width  int
	height int
}

func (e *ErrInvalidSizeVo) Error() string {
	return fmt.Sprintf("width (%d) and height(%d) of size must be greater than 0", e.width, e.height)
}

type SizeVo struct {
	width  int
	height int
}

func NewSizeVo(width int, height int) (SizeVo, error) {
	if width < 1 || height < 1 {
		return SizeVo{}, &ErrInvalidSizeVo{width: width, height: height}
	}

	return SizeVo{
		width:  width,
		height: height,
	}, nil
}

func (size SizeVo) GetWidth() int {
	return size.width
}

func (size SizeVo) GetHeight() int {
	return size.height
}

func (size SizeVo) CoverLocation(location LocationVo) bool {
	includesAll := true
	if location.x >= size.width || location.y >= size.height {
		includesAll = false
	}

	return includesAll
}
