package playermodel

type PlayerAgr struct {
	id   PlayerIdVo
	name string
}

func NewPlayerAgr(id PlayerIdVo, name string) PlayerAgr {
	return PlayerAgr{
		id:   id,
		name: name,
	}
}

func (p PlayerAgr) GetId() PlayerIdVo {
	return p.id
}

func (p PlayerAgr) GetName() string {
	return p.name
}
