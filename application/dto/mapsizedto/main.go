package mapsizedto

import "github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"

type DTO struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func ToDTO(mapSize valueobject.MapSize) DTO {
	return DTO{
		Width:  mapSize.GetWidth(),
		Height: mapSize.GetHeight(),
	}
}
