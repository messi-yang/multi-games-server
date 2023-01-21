package livegamemodel

type PlayerEntity struct {
	id   PlayerIdVo
	name string
}

func NewPlayerEntity(id PlayerIdVo, name string) PlayerEntity {
	return PlayerEntity{
		id:   id,
		name: name,
	}
}

func (p PlayerEntity) GetId() PlayerIdVo {
	return p.id
}

func (p PlayerEntity) GetName() string {
	return p.name
}
