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

func (ms MapSize) IncludesMapRange(mapRange MapRange) bool {
	if mapRange.from.x < 0 || mapRange.from.x >= ms.width {
		return false
	}
	if mapRange.to.x < 0 || mapRange.to.x >= ms.width {
		return false
	}
	if mapRange.from.y < 0 || mapRange.from.y >= ms.height {
		return false
	}
	if mapRange.to.y < 0 || mapRange.to.y >= ms.height {
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

func (ms MapSize) IncludesAllLocations(locations []Location) bool {
	includesAll := true
	for _, location := range locations {
		if location.x >= ms.width || location.y >= ms.height {
			includesAll = false
		}
	}

	return includesAll
}
