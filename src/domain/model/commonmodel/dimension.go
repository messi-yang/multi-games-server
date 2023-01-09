package commonmodel

import (
	"fmt"
)

type ErrInvalidDimension struct {
	width  int
	height int
}

func (e *ErrInvalidDimension) Error() string {
	return fmt.Sprintf("width (%d) and height(%d) of dimension must be greater than 0", e.width, e.height)
}

type Dimension struct {
	width  int
	height int
}

func NewDimension(width int, height int) (Dimension, error) {
	if width < 1 || height < 1 {
		return Dimension{}, &ErrInvalidDimension{width: width, height: height}
	}

	return Dimension{
		width:  width,
		height: height,
	}, nil
}

func (ms Dimension) GetWidth() int {
	return ms.width
}

func (ms Dimension) GetHeight() int {
	return ms.height
}

func (ms Dimension) IncludesRange(rangeVo Range) bool {
	if rangeVo.from.x < 0 || rangeVo.from.x >= ms.width {
		return false
	}
	if rangeVo.to.x < 0 || rangeVo.to.x >= ms.width {
		return false
	}
	if rangeVo.from.y < 0 || rangeVo.from.y >= ms.height {
		return false
	}
	if rangeVo.to.y < 0 || rangeVo.to.y >= ms.height {
		return false
	}
	return true
}

func (ms Dimension) IncludesLocation(location Location) bool {
	includesAll := true
	if location.x >= ms.width || location.y >= ms.height {
		includesAll = false
	}

	return includesAll
}
