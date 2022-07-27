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

func (gr *GameRoom) AreCoordinatesValid(coordinates []valueobject.Coordinate) bool {
	for _, coord := range coordinates {
		x := coord.GetX()
		y := coord.GetY()
		if x < 0 || x >= gr.game.MapSize.GetWidth() {
			return false
		}
		if y < 0 || y >= gr.game.MapSize.GetHeight() {
			return false
		}
	}
	return true
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

func (gr *GameRoom) GetUnitMap() [][]valueobject.Unit {
	return gr.game.UnitMap
}

func (gr *GameRoom) UpdateUnitMap(unitMap [][]valueobject.Unit) error {
	gr.game.UnitMap = unitMap
	return nil
}

func (gr *GameRoom) GetUnitMapWithArea(area valueobject.Area) ([][]valueobject.Unit, error) {
	width := area.GetTo().GetX() - area.GetFrom().GetX() + 1
	height := area.GetTo().GetY() - area.GetFrom().GetY() + 1
	gameMap := make([][]valueobject.Unit, width)
	for x := 0; x < width; x += 1 {
		gameMap[x] = make([]valueobject.Unit, height)
		for y := 0; y < height; y += 1 {
			gameMap[x][y] = gr.game.UnitMap[x][y]
		}
	}

	return gameMap, nil
}

func (gr *GameRoom) GetUnitsWithCoordinates(coordinates []valueobject.Coordinate) ([]valueobject.Unit, error) {
	units := make([]valueobject.Unit, 0)
	for _, coord := range coordinates {
		unit := gr.game.UnitMap[coord.GetX()][coord.GetY()]
		units = append(units, unit)
	}

	return units, nil
}

func (gr *GameRoom) GetUnit(coordinate valueobject.Coordinate) valueobject.Unit {
	return gr.game.UnitMap[coordinate.GetX()][coordinate.GetY()]
}

func (gr *GameRoom) UpdateUnit(coordinate valueobject.Coordinate, unit valueobject.Unit) error {
	gr.game.UnitMap[coordinate.GetX()][coordinate.GetY()] = unit

	return nil
}
