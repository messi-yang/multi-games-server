package mapsizedto

import "github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"

type Dto struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func ToDto(mapSize valueobject.MapSize) Dto {
	return Dto{
		Width:  mapSize.GetWidth(),
		Height: mapSize.GetHeight(),
	}
}
