package livegamemodel

import "github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"

type Camera struct {
	center commonmodel.Location
}

func NewCamera(center commonmodel.Location) Camera {
	return Camera{
		center: center,
	}
}

func (camera Camera) GetCenter() commonmodel.Location {
	return camera.center
}
