package valueobject

import (
	"fmt"
)

type ErrInvalidMapSize struct {
	width  int
	height int
}

func (e *ErrInvalidMapSize) Error() string {
	return fmt.Sprintf("width (%d) and height(%d) of map size must be greater than 0", e.width, e.height)
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

func (ms MapSize) CoversArea(area Area) bool {
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
