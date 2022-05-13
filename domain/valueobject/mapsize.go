package valueobject

type MapSize struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func NewMapSize(width int, height int) MapSize {
	return MapSize{
		Width:  width,
		Height: height,
	}
}

func (ms MapSize) SetWidth(width int) MapSize {
	ms.Width = width

	return ms
}

func (ms MapSize) SetHeight(height int) MapSize {
	ms.Height = height

	return ms
}
