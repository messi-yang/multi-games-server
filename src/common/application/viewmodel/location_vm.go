package viewmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"

type LocationVm struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func NewLocationVm(location commonmodel.LocationVo) LocationVm {
	return LocationVm{
		X: location.GetX(),
		Y: location.GetY(),
	}
}
