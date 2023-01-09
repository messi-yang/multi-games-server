package playermodel

type Player struct {
	id   PlayerId
	name string
}

func NewPlayer(id PlayerId, name string) Player {
	return Player{
		id:   id,
		name: name,
	}
}

func (p Player) GetId() PlayerId {
	return p.id
}

func (p Player) GetName() string {
	return p.name
}
