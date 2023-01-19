package playermodel

type PlayerAgg struct {
	id   PlayerIdVo
	name string
}

func NewPlayerAgg(id PlayerIdVo, name string) PlayerAgg {
	return PlayerAgg{
		id:   id,
		name: name,
	}
}

func (p PlayerAgg) GetId() PlayerIdVo {
	return p.id
}

func (p PlayerAgg) GetName() string {
	return p.name
}
