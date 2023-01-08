package viewmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"

type MapSize struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func NewMapSize(mapSize commonmodel.MapSize) MapSize {
	return MapSize{
		Width:  mapSize.GetWidth(),
		Height: mapSize.GetHeight(),
	}
}

func (dto MapSize) ToValueObject() (commonmodel.MapSize, error) {
	return commonmodel.NewMapSize(dto.Width, dto.Height)
}
