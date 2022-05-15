package aggregate

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/entity"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/valueobject"
	"github.com/google/uuid"
)

type GameRoom struct {
	game *entity.Game
}

func NewGameRoom() GameRoom {
	newGame := entity.NewGame()
	return GameRoom{
		game: &newGame,
	}
}

func (gr *GameRoom) GetGameId() uuid.UUID {
	game := gr.game
	return game.Id
}

func (gr *GameRoom) GetGameMapSize() valueobject.MapSize {
	game := gr.game
	return game.MapSize
}

func (gr *GameRoom) UpdateGameMapSize(mapSize valueobject.MapSize) error {
	gr.game.MapSize = mapSize
	return nil
}

func (gr *GameRoom) GetGameUnitMatrix() [][]valueobject.GameUnit {
	game := gr.game
	return game.UnitMatrix
}

func (gr *GameRoom) UpdateGameUnitMatrix(unitMatrix [][]valueobject.GameUnit) error {
	gr.game.UnitMatrix = unitMatrix
	return nil
}
