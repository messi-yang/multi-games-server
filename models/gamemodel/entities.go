package gamemodel

type GameFieldEntity [][]GameFieldUnitEntity

type GameFieldUnitEntity struct {
	Alive bool `json:"alive"`
	Age   int  `json:"age"`
}

type GameFieldSizeEntity struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}
