package mapsizedto

import "github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"

type MapSizeDTO struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func ToDTO(mapSize valueobject.MapSize) MapSizeDTO {
	return MapSizeDTO{
		Width:  mapSize.GetWidth(),
		Height: mapSize.GetHeight(),
	}
}
