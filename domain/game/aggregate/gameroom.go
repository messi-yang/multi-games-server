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
	if !gr.isMapSizeValid(mapSize) {
		return ErrInvalidMapSize
	}
	gr.game.MapSize = mapSize
	return nil
}

func (gr *GameRoom) GetUnitMap() [][]valueobject.Unit {
	return gr.game.UnitMap
}

func (gr *GameRoom) UpdateUnitMap(unitMap [][]valueobject.Unit) {
	gr.game.UnitMap = unitMap
}

func (gr *GameRoom) GetUnitMapWithArea(area valueobject.Area) ([][]valueobject.Unit, error) {
	if !gr.isAreaValid(area) {
		return nil, ErrInvalidArea
	}
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

func (gr *GameRoom) UpdateUnit(coordinate valueobject.Coordinate, unit valueobject.Unit) {
	adjustedCoordinate := gr.adjustCoordinate(coordinate)
	gr.game.UnitMap[adjustedCoordinate.GetX()][adjustedCoordinate.GetY()] = unit
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

func (gr *GameRoom) isMapSizeValid(mapSize valueobject.MapSize) bool {
	if mapSize.GetHeight() <= 0 || mapSize.GetWidth() <= 0 {
		return false
	}
	return true
}

func (gr *GameRoom) isCoordinateValid(coord valueobject.Coordinate) bool {
	if coord.GetX() < 0 || coord.GetX() >= gr.GetGameMapSize().GetWidth() {
		return false
	}
	if coord.GetY() < 0 || coord.GetY() >= gr.GetGameMapSize().GetHeight() {
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

	if areaWidth > gr.GetGameMapSize().GetWidth() || areaHeght > gr.GetGameMapSize().GetHeight() {
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
