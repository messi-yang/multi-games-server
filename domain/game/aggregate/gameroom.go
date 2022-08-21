package aggregate

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/entity"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"
	"github.com/google/uuid"
)

var (
	ErrAreaExceedsUnitMap = errors.New("area should contain valid from and to coordinates and it should never exceed map size")
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
	if !gr.GetUnitMapSize().IncludesArea(area) {
		return nil, ErrAreaExceedsUnitMap
	}
	offsetX := area.GetFrom().GetX()
	offsetY := area.GetFrom().GetY()
	areaWidth := area.GetWidth()
	areaHeight := area.GetHeight()
	gameMap := make(valueobject.UnitMap, areaWidth)
	for x := 0; x < areaWidth; x += 1 {
		gameMap[x] = make([]valueobject.Unit, areaHeight)
		for y := 0; y < areaHeight; y += 1 {
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

func (gr *GameRoom) UpdateUnit(coordinate valueobject.Coordinate, unit valueobject.Unit) {
	adjustedCoordinate := gr.adjustCoordinate(coordinate)
	gr.game.SetUnit(adjustedCoordinate, unit)
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
