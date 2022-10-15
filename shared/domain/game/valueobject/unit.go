package valueobject

type Unit struct {
	alive bool
}

func NewUnit(alive bool) Unit {
	return Unit{
		alive: alive,
	}
}

func (gu Unit) GetAlive() bool {
	return gu.alive
}

func (gu Unit) SetAlive(alive bool) Unit {
	gu.alive = alive
	return gu
}
