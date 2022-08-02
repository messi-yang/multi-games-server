package aggregate

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/entity"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"
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

func (gr *GameRoom) GetUnitMap() [][]valueobject.Unit {
	return gr.game.UnitMap
}

func (gr *GameRoom) UpdateUnitMap(unitMap [][]valueobject.Unit) error {
	gr.game.UnitMap = unitMap
	return nil
}

func (gr *GameRoom) GetUnitMapWithArea(area valueobject.Area) ([][]valueobject.Unit, error) {
	offsetX := area.GetFrom().GetX()
	offsetY := area.GetFrom().GetY()
	width := area.GetTo().GetX() - area.GetFrom().GetX() + 1
	height := area.GetTo().GetY() - area.GetFrom().GetY() + 1
	gameMap := make([][]valueobject.Unit, width)
	for x := 0; x < width; x += 1 {
		gameMap[x] = make([]valueobject.Unit, height)
		for y := 0; y < height; y += 1 {
			gameMap[x][y] = gr.game.UnitMap[x+offsetX][y+offsetY]
		}
	}

	return gameMap, nil
}

func (gr *GameRoom) GetUnitsWithCoordinates(coordinates []valueobject.Coordinate) ([]valueobject.Unit, error) {
	units := make([]valueobject.Unit, 0)
	for _, coord := range coordinates {
		adjustedCoordinate := gr.adjustCoordinate(coord)
		unit := gr.game.UnitMap[adjustedCoordinate.GetX()][adjustedCoordinate.GetY()]
		units = append(units, unit)
	}

	return units, nil
}

func (gr *GameRoom) GetUnit(coordinate valueobject.Coordinate) valueobject.Unit {
	adjustedCoordinate := gr.adjustCoordinate(coordinate)
	return gr.game.UnitMap[adjustedCoordinate.GetX()][adjustedCoordinate.GetY()]
}

func (gr *GameRoom) UpdateUnit(coordinate valueobject.Coordinate, unit valueobject.Unit) error {
	adjustedCoordinate := gr.adjustCoordinate(coordinate)
	gr.game.UnitMap[adjustedCoordinate.GetX()][adjustedCoordinate.GetY()] = unit

	return nil
}

func (gr *GameRoom) GetCoordinatesInArea(coordinates []valueobject.Coordinate, area valueobject.Area) []valueobject.Coordinate {
	coordinatesInArea := make([]valueobject.Coordinate, 0)
	for _, coordinate := range coordinates {
		if gr.isCoordinateInArea(coordinate, area) {
			coordinatesInArea = append(coordinatesInArea, coordinate)
		}
	}

	return coordinatesInArea
}

func (gr *GameRoom) isCoordinateInArea(coordinate valueobject.Coordinate, area valueobject.Area) bool {
	x := coordinate.GetX()
	y := coordinate.GetY()
	if x < area.GetFrom().GetX() || x > area.GetTo().GetX() {
		return false
	}
	if y < area.GetFrom().GetY() || y > area.GetTo().GetY() {
		return false
	}
	return true
}

func (gr *GameRoom) adjustCoordinate(coordinate valueobject.Coordinate) valueobject.Coordinate {
	adjustedX := coordinate.GetX()
	adjustedY := coordinate.GetY()
	mapWidth := gr.GetGameMapSize().GetWidth()
	mapHeight := gr.GetGameMapSize().GetHeight()

	for adjustedX < 0 {
		adjustedX += mapWidth
	}

	for adjustedY < 0 {
		adjustedY += mapHeight
	}

	adjustedCoordinate := valueobject.NewCoordinate(adjustedX%mapWidth, adjustedY%mapHeight)

	return adjustedCoordinate
}
