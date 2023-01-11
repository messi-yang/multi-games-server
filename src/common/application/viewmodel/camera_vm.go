package viewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/livegamemodel"
)

type CameraVm struct {
	Center LocationVm `json:"center"`
}

func NewCameraVm(camera livegamemodel.Camera) CameraVm {
	return CameraVm{
		Center: NewLocationVm(camera.GetCenter()),
	}
}

func (camera CameraVm) ToValueObject() (livegamemodel.Camera, error) {
	location, err := commonmodel.NewLocation(camera.Center.X, camera.Center.Y)
	if err != nil {
		return livegamemodel.Camera{}, err
	}
	return livegamemodel.NewCamera(location), nil
}
