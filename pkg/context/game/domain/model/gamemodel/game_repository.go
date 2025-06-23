package gamemodel

type GameRepo interface {
	Add(Game) error
	Get(GameId) (Game, error)
	Update(Game) error
}
