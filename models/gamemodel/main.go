package gamemodel

type GameField [][]GameFieldUnit

type GameFieldUnit struct {
	Alive bool `json:"alive"`
	Age   int  `json:"age"`
}

type GameFieldSize struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}
