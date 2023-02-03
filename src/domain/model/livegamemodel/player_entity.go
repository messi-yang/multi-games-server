package livegamemodel

import "github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"

type PlayerEntity struct {
	id       PlayerIdVo
	name     string
	camera   CameraVo
	location commonmodel.LocationVo
}

func NewPlayerEntity(id PlayerIdVo, name string, camera CameraVo, location commonmodel.LocationVo) PlayerEntity {
	return PlayerEntity{
		id:       id,
		name:     name,
		camera:   camera,
		location: location,
	}
}

func (p *PlayerEntity) GetId() PlayerIdVo {
	return p.id
}

func (p *PlayerEntity) GetName() string {
	return p.name
}

func (p *PlayerEntity) GetCamera() CameraVo {
	return p.camera
}

func (p *PlayerEntity) ChangeCamera(camera CameraVo) {
	p.camera = camera
}

func (p *PlayerEntity) GetLocation() commonmodel.LocationVo {
	return p.location
}

func (p *PlayerEntity) ChangeLocation(location commonmodel.LocationVo) {
	p.location = location
}
