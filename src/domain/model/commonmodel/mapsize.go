package commonmodel

import (
	"fmt"
)

type ErrInvalidMapSize struct {
	width  int
	height int
}

func (e *ErrInvalidMapSize) Error() string {
	return fmt.Sprintf("width (%d) and height(%d) of mapSize must be greater than 0", e.width, e.height)
}

type MapSize struct {
	width  int
	height int
}

func NewMapSize(width int, height int) (MapSize, error) {
	if width < 1 || height < 1 {
		return MapSize{}, &ErrInvalidMapSize{width: width, height: height}
	}

	return MapSize{
		width:  width,
		height: height,
	}, nil
}

func (ms MapSize) GetWidth() int {
	return ms.width
}

func (ms MapSize) GetHeight() int {
	return ms.height
}

func (ms MapSize) IncludesExtent(extent Extent) bool {
	if extent.from.x < 0 || extent.from.x >= ms.width {
		return false
	}
	if extent.to.x < 0 || extent.to.x >= ms.width {
		return false
	}
	if extent.from.y < 0 || extent.from.y >= ms.height {
		return false
	}
	if extent.to.y < 0 || extent.to.y >= ms.height {
		return false
	}
	return true
}

func (ms MapSize) IncludesLocation(location Location) bool {
	includesAll := true
	if location.x >= ms.width || location.y >= ms.height {
		includesAll = false
	}

	return includesAll
}
