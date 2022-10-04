package dto

import "github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"

type MapSizeDto struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func NewMapSizeDto(mapSize valueobject.MapSize) MapSizeDto {
	return MapSizeDto{
		Width:  mapSize.GetWidth(),
		Height: mapSize.GetHeight(),
	}
}
