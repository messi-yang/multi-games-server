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

func (gr *GameRoom) GetGame() entity.Game {
	return *gr.game
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

func (gr *GameRoom) GetUnitMatrix() [][]valueobject.Unit {
	return gr.game.UnitMatrix
}

func (gr *GameRoom) UpdateUnitMatrix(unitMatrix [][]valueobject.Unit) error {
	gr.game.UnitMatrix = unitMatrix
	return nil
}

func (gr *GameRoom) GetUnitMatrixWithArea(area valueobject.Area) ([][]valueobject.Unit, error) {
	width := area.GetTo().GetX() - area.GetFrom().GetX() + 1
	height := area.GetTo().GetY() - area.GetFrom().GetY() + 1
	gameMatrix := make([][]valueobject.Unit, width)
	for x := 0; x < width; x += 1 {
		gameMatrix[x] = make([]valueobject.Unit, height)
		for y := 0; y < height; y += 1 {
			gameMatrix[x][y] = gr.game.UnitMatrix[x][y]
		}
	}

	return gameMatrix, nil
}

func (gr *GameRoom) GetUnitsWithCoordinates(coordinates []valueobject.Coordinate) ([]valueobject.Unit, error) {
	units := make([]valueobject.Unit, 0)
	for _, coord := range coordinates {
		unit := gr.game.UnitMatrix[coord.GetX()][coord.GetY()]
		units = append(units, unit)
	}

	return units, nil
}

func (gr *GameRoom) GetUnit(coordinate valueobject.Coordinate) valueobject.Unit {
	return gr.game.UnitMatrix[coordinate.GetX()][coordinate.GetY()]
}

func (gr *GameRoom) UpdateUnit(coordinate valueobject.Coordinate, unit valueobject.Unit) error {
	gr.game.UnitMatrix[coordinate.GetX()][coordinate.GetY()] = unit

	return nil
}
