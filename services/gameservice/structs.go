package gameservice

type GameUnits [][]GameUnit

type GameUnit struct {
	Alive bool `json:"alive"`
	Age   int  `json:"age"`
}

type GameSize struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type GameCoordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type GameArea struct {
	From GameCoordinate `json:"from"`
	To   GameCoordinate `json:"to"`
}
