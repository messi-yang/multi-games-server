package gamedao

type GameUnits [][]GameUnit

type GameUnit struct {
	Alive bool `json:"alive"`
	Age   int  `json:"age"`
}

type GameSize struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}
