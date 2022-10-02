package valueobject

type Unit struct {
	alive bool
	age   int
}

func NewUnit(alive bool, age int) Unit {
	return Unit{
		alive: alive,
		age:   age,
	}
}

func (gu Unit) GetAlive() bool {
	return gu.alive
}

func (gu Unit) GetAge() int {
	return gu.age
}

func (gu Unit) SetAlive(alive bool) Unit {
	gu.alive = alive
	return gu
}
