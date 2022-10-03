package entity

import "github.com/google/uuid"

type Player struct {
	id uuid.UUID
}

func NewPlayer() Player {
	return Player{id: uuid.New()}
}

func NewPlayerWithExistingId(id uuid.UUID) Player {
	return Player{id: id}
}

func (p *Player) GetId() uuid.UUID {
	return p.id
}
