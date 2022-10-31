package valueobject

import "github.com/google/uuid"

type Unit struct {
	alive bool
	id    uuid.UUID
}

func NewUnit(alive bool, uuid uuid.UUID) Unit {
	return Unit{
		alive: alive,
		id:    uuid,
	}
}

func (gu Unit) GetAlive() bool {
	return gu.alive
}

func (gu Unit) SetAlive(alive bool) Unit {
	return Unit{
		alive: alive,
		id:    gu.id,
	}
}

func (gu Unit) GetId() uuid.UUID {
	return gu.id
}

func (gu Unit) SetId(uuid uuid.UUID) Unit {
	return Unit{
		alive: gu.alive,
		id:    uuid,
	}
}
