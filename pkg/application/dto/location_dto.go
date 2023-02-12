package dto

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"

type LocationDto struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func NewLocationDto(location commonmodel.LocationVo) LocationDto {
	return LocationDto{
		X: location.GetX(),
		Y: location.GetY(),
	}
}
