package valueobject

type GameUnit struct {
	Alive bool `json:"alive"`
	Age   int  `json:"age"`
}

func NewGameUnit(alive bool, age int) GameUnit {
	return GameUnit{
		Alive: alive,
		Age:   age,
	}
}

func (gu GameUnit) SetAlive(alive bool) GameUnit {
	gu.Alive = alive
	return gu
}

func (gu GameUnit) SetAge(age int) GameUnit {
	gu.Age = age
	return gu
}
