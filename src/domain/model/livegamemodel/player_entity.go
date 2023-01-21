package livegamemodel

type PlayerEntity struct {
	id     PlayerIdVo
	name   string
	camera CameraVo
}

func NewPlayerEntity(id PlayerIdVo, name string, camera CameraVo) PlayerEntity {
	return PlayerEntity{
		id:     id,
		name:   name,
		camera: camera,
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
