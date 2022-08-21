package aggregate

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/entity"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"
	"github.com/google/uuid"
)

var (
	ErrInvalidMapSize = errors.New("width or height of map size cannot be smaller than 1")
	ErrInvalidArea    = errors.New("area should contain valid from and to coordinates and it should never exceed map size")
)

type GameRoom struct {
	game *entity.Game
}

func NewGameRoom(game entity.Game) GameRoom {
	return GameRoom{
		game: &game,
	}
}

func (gr *GameRoom) GetGameId() uuid.UUID {
	return gr.game.GetId()
}

func (gr *GameRoom) GetUnitMapSize() valueobject.MapSize {
	return gr.game.GetUnitMapSize()
}

func (gr *GameRoom) GetUnitMap() valueobject.UnitMap {
	return gr.game.GetUnitMap()
}

func (gr *GameRoom) UpdateUnitMap(unitMap valueobject.UnitMap) {
	gr.game.SetUnitMap(unitMap)
}

func (gr *GameRoom) GetUnitMapWithArea(area valueobject.Area) (valueobject.UnitMap, error) {
	if !gr.isAreaValid(area) {
		return nil, ErrInvalidArea
	}
	offsetX := area.GetFrom().GetX()
	offsetY := area.GetFrom().GetY()
	width := area.GetTo().GetX() - area.GetFrom().GetX() + 1
	height := area.GetTo().GetY() - area.GetFrom().GetY() + 1
	gameMap := make(valueobject.UnitMap, width)
	for x := 0; x < width; x += 1 {
		gameMap[x] = make([]valueobject.Unit, height)
		for y := 0; y < height; y += 1 {
			coordinate := valueobject.NewCoordinate(x+offsetX, y+offsetY)
			gameMap[x][y] = gr.game.GetUnit(coordinate)
		}
	}

	return gameMap, nil
}

func (gr *GameRoom) GetUnitsWithCoordinates(coordinates []valueobject.Coordinate) ([]valueobject.Unit, error) {
	units := make([]valueobject.Unit, 0)
	for _, coord := range coordinates {
		adjustedCoordinate := gr.adjustCoordinate(coord)
		unit := gr.game.GetUnit(adjustedCoordinate)
		units = append(units, unit)
	}

	return units, nil
}

func (gr *GameRoom) GetUnit(coordinate valueobject.Coordinate) valueobject.Unit {
	adjustedCoordinate := gr.adjustCoordinate(coordinate)
	return gr.game.GetUnit(adjustedCoordinate)
}

func (gr *GameRoom) UpdateUnit(coordinate valueobject.Coordinate, unit valueobject.Unit) {
	adjustedCoordinate := gr.adjustCoordinate(coordinate)
	gr.game.SetUnit(adjustedCoordinate, unit)
}

func (gr *GameRoom) FilterCoordinatesWithArea(coordinates []valueobject.Coordinate, area valueobject.Area) ([]valueobject.Coordinate, error) {
	if !gr.isAreaValid(area) {
		return nil, ErrInvalidArea
	}
	coordinatesInArea := make([]valueobject.Coordinate, 0)
	for _, coordinate := range coordinates {
		if gr.isCoordinateInArea(coordinate, area) {
			coordinatesInArea = append(coordinatesInArea, coordinate)
		}
	}

	return coordinatesInArea, nil
}

func (gr *GameRoom) isCoordinateValid(coord valueobject.Coordinate) bool {
	if coord.GetX() < 0 || coord.GetX() >= gr.GetUnitMapSize().GetWidth() {
		return false
	}
	if coord.GetY() < 0 || coord.GetY() >= gr.GetUnitMapSize().GetHeight() {
		return false
	}
	return true
}

func (gr *GameRoom) isAreaValid(area valueobject.Area) bool {
	if !gr.isCoordinateValid(area.GetFrom()) {
		return false
	}
	if !gr.isCoordinateValid(area.GetTo()) {
		return false
	}
	if area.GetFrom().GetX() > area.GetTo().GetX() {
		return false
	}
	if area.GetFrom().GetY() > area.GetTo().GetY() {
		return false
	}

	areaWidth := area.GetTo().GetX() - area.GetFrom().GetX()
	areaHeght := area.GetTo().GetY() - area.GetFrom().GetY()

	if areaWidth > gr.GetUnitMapSize().GetWidth() || areaHeght > gr.GetUnitMapSize().GetHeight() {
		return false
	}
	return true
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
	mapWidth := gr.GetUnitMapSize().GetWidth()
	mapHeight := gr.GetUnitMapSize().GetHeight()

	for adjustedX < 0 {
		adjustedX += mapWidth
	}

	for adjustedY < 0 {
		adjustedY += mapHeight
	}

	adjustedCoordinate := valueobject.NewCoordinate(adjustedX%mapWidth, adjustedY%mapHeight)

	return adjustedCoordinate
}
