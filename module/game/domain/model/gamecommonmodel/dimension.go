package gamecommonmodel

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

func (ms Dimension) IncludesArea(area Area) bool {
	if area.from.x < 0 || area.from.x >= ms.width {
		return false
	}
	if area.to.x < 0 || area.to.x >= ms.width {
		return false
	}
	if area.from.y < 0 || area.from.y >= ms.height {
		return false
	}
	if area.to.y < 0 || area.to.y >= ms.height {
		return false
	}
	return true
}

func (ms Dimension) IncludesAllCoordinates(coordinates []Coordinate) bool {
	includesAll := true
	for _, coordinate := range coordinates {
		if coordinate.x >= ms.width || coordinate.y >= ms.height {
			includesAll = false
		}
	}

	return includesAll
}
