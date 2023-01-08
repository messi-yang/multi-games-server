package viewmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"

type MapSizeVm struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func NewMapSizeVm(mapSize commonmodel.MapSize) MapSizeVm {
	return MapSizeVm{
		Width:  mapSize.GetWidth(),
		Height: mapSize.GetHeight(),
	}
}

func (dto MapSizeVm) ToValueObject() (commonmodel.MapSize, error) {
	return commonmodel.NewMapSize(dto.Width, dto.Height)
}
