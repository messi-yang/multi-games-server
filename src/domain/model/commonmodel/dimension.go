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

func (ms Dimension) IncludesBound(bound_ Bound) bool {
	if bound_.from.x < 0 || bound_.from.x >= ms.width {
		return false
	}
	if bound_.to.x < 0 || bound_.to.x >= ms.width {
		return false
	}
	if bound_.from.y < 0 || bound_.from.y >= ms.height {
		return false
	}
	if bound_.to.y < 0 || bound_.to.y >= ms.height {
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
