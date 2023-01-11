package viewmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/livegamemodel"

type CameraVm struct {
	Center LocationVm `json:"center"`
}

func NewCameraVm(camera livegamemodel.Camera) CameraVm {
	return CameraVm{
		Center: NewLocationVm(camera.GetCenter()),
	}
}
