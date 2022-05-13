package aggregate

import "github.com/DumDumGeniuss/game-of-liberty-computer/domain/entity"

type GameRoom struct {
	Game *entity.Game
}

func NewGameRoom(game *entity.Game) GameRoom {
	return GameRoom{
		Game: game,
	}
}
