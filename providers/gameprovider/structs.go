package gameprovider

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
	X int
	Y int
}

type GameArea struct {
	From GameCoordinate
	To   GameCoordinate
}
