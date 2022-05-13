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
