package aggregate

import "github.com/google/uuid"

type Player struct {
	id uuid.UUID
}

func NewPlayer(id uuid.UUID) Player {
	return Player{id: uuid.New()}
}

func (p *Player) GetId() uuid.UUID {
	return p.id
}
