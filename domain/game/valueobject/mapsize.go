package valueobject

type MapSize struct {
	width  int
	height int
}

func NewMapSize(width int, height int) MapSize {
	return MapSize{
		width:  width,
		height: height,
	}
}

func (ms MapSize) GetWidth() int {
	return ms.width
}

func (ms MapSize) GetHeight() int {
	return ms.height
}
