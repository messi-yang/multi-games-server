package gameentity

type GameUnit struct {
	Alive bool `json:"alive"`
	Age   int  `json:"age"`
}

type gameUnits [][]*GameUnit

type gameBlockChangeEventCallback func([][]*GameUnit)

type gameBlockChangeEventListner struct {
	fromX            int
	fromY            int
	toX              int
	toY              int
	gameUnitsChannel chan gameUnits
	stopChannel      chan bool
}
