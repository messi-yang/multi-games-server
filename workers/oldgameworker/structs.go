package oldgameworker

type GameUnit struct {
	Alive bool `json:"alive"`
	Age   int  `json:"age"`
}

type gameUnits [][]*GameUnit

type gameBlockChangeEventCallback func([][]*GameUnit)

type GameBlockArea struct {
	FromX int
	FromY int
	ToX   int
	ToY   int
}

type gameBlockChangeEventSubscriber struct {
	key              string
	gameBlockArea    GameBlockArea
	gameUnitsChannel chan gameUnits
}
