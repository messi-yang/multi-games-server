package gamemodel

type GameUnitsModel [][]GameUnitModel

type GameUnitModel struct {
	Alive bool `json:"alive"`
	Age   int  `json:"age"`
}

type GameSizeModel struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}
