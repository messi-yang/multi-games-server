package dto

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"

type LocationDto struct {
	X int `json:"x"`
	Z int `json:"z"`
}

func NewLocationDto(location commonmodel.LocationVo) LocationDto {
	return LocationDto{
		X: location.GetX(),
		Z: location.GetZ(),
	}
}
