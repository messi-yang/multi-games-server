package aggregate

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/entity"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/valueobject"
	"github.com/google/uuid"
)

type GameRoom struct {
	game entity.Game
}

func NewGameRoom(game entity.Game) GameRoom {
	return GameRoom{
		game: game,
	}
}

func (gr *GameRoom) GetGame() entity.Game {
	return gr.game
}

func (gr *GameRoom) GetGameId() uuid.UUID {
	game := gr.game
	return game.GetId()
}

func (gr *GameRoom) GetGameMapSize() valueobject.MapSize {
	game := gr.game
	return game.GetMapSize()
}

func (gr *GameRoom) GetGameUnitMatrix() [][]valueobject.GameUnit {
	game := gr.game
	return game.GetUnitMatrix()
}
