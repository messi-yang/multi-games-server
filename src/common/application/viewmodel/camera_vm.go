package viewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/livegamemodel"
)

type CameraVm struct {
	Center LocationVm `json:"center"`
}

func NewCameraVm(camera livegamemodel.CameraVo) CameraVm {
	return CameraVm{
		Center: NewLocationVm(camera.GetCenter()),
	}
}

func (camera CameraVm) ToValueObject() (livegamemodel.CameraVo, error) {
	location, err := commonmodel.NewLocationVo(camera.Center.X, camera.Center.Y)
	if err != nil {
		return livegamemodel.CameraVo{}, err
	}
	return livegamemodel.NewCameraVo(location), nil
}
