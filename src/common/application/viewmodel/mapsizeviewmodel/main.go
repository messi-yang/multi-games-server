package mapsizeviewmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"

type ViewModel struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func New(mapSize commonmodel.MapSize) ViewModel {
	return ViewModel{
		Width:  mapSize.GetWidth(),
		Height: mapSize.GetHeight(),
	}
}

func (dto ViewModel) ToValueObject() (commonmodel.MapSize, error) {
	return commonmodel.NewMapSize(dto.Width, dto.Height)
}
