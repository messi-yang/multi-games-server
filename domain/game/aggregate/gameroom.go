package aggregate

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/entity"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/valueobject"
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
	return gr.game.Id
}

func (gr *GameRoom) GetGameMapSize() valueobject.MapSize {
	return gr.game.MapSize
}

func (gr *GameRoom) UpdateGameMapSize(mapSize valueobject.MapSize) error {
	gr.game.MapSize = mapSize
	return nil
}

func (gr *GameRoom) GetGameUnitMatrix() [][]valueobject.GameUnit {
	return gr.game.UnitMatrix
}

func (gr *GameRoom) UpdateGameUnitMatrix(unitMatrix [][]valueobject.GameUnit) error {
	gr.game.UnitMatrix = unitMatrix
	return nil
}

func (gr *GameRoom) GetGameUnitMatrixWithArea(area valueobject.Area) ([][]valueobject.GameUnit, error) {
	width := area.GetTo().GetX() - area.GetFrom().GetX() + 1
	height := area.GetTo().GetY() - area.GetFrom().GetY() + 1
	gameMatrix := make([][]valueobject.GameUnit, width)
	for x := 0; x < width; x += 1 {
		gameMatrix[x] = make([]valueobject.GameUnit, height)
		for y := 0; y < height; y += 1 {
			gameMatrix[x][y] = gr.game.UnitMatrix[x][y]
		}
	}

	return gameMatrix, nil
}

func (gr *GameRoom) GetGameUnit(coordinate valueobject.Coordinate) valueobject.GameUnit {
	return gr.game.UnitMatrix[coordinate.GetX()][coordinate.GetY()]
}

func (gr *GameRoom) UpdateGameUnit(coordinate valueobject.Coordinate, gameUnit valueobject.GameUnit) error {
	gr.game.UnitMatrix[coordinate.GetX()][coordinate.GetY()] = gameUnit

	return nil
}
