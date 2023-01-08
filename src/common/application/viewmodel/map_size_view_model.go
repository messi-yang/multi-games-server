package viewmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"

type MapSizeViewModel struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func NewMapSizeViewModel(mapSize commonmodel.MapSize) MapSizeViewModel {
	return MapSizeViewModel{
		Width:  mapSize.GetWidth(),
		Height: mapSize.GetHeight(),
	}
}

func (dto MapSizeViewModel) ToValueObject() (commonmodel.MapSize, error) {
	return commonmodel.NewMapSize(dto.Width, dto.Height)
}
