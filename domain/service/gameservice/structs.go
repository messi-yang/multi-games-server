package gameservice

type GameCoordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type GameArea struct {
	From GameCoordinate `json:"from"`
	To   GameCoordinate `json:"to"`
}
