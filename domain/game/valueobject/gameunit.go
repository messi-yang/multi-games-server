package valueobject

type GameUnit struct {
	alive bool
	age   int
}

func NewGameUnit(alive bool, age int) GameUnit {
	return GameUnit{
		alive: alive,
		age:   age,
	}
}

func (gu GameUnit) GetAlive() bool {
	return gu.alive
}

func (gu GameUnit) GetAge() int {
	return gu.age
}
