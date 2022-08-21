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
