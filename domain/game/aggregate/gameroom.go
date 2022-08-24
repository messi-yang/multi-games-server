package aggregate

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/entity"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"
	"github.com/google/uuid"
)

var (
	ErrAreaExceedsUnitMap              = errors.New("area should contain valid from and to coordinates and it should never exceed map size")
	ErrSomeCoordinatesNotIncludedInMap = errors.New("some coordinates are not included in the unit map")
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

func (gr *GameRoom) GetUnitMapByArea(area valueobject.Area) (valueobject.UnitMap, error) {
	if !gr.GetUnitMapSize().IncludesArea(area) {
		return valueobject.UnitMap{}, ErrAreaExceedsUnitMap
	}
	offsetX := area.GetFrom().GetX()
	offsetY := area.GetFrom().GetY()
	areaWidth := area.GetWidth()
	areaHeight := area.GetHeight()
	unitMatrix := make([][]valueobject.Unit, areaWidth)
	for x := 0; x < areaWidth; x += 1 {
		unitMatrix[x] = make([]valueobject.Unit, areaHeight)
		for y := 0; y < areaHeight; y += 1 {
			coordinate, _ := valueobject.NewCoordinate(x+offsetX, y+offsetY)
			unitMatrix[x][y] = gr.game.GetUnit(coordinate)
		}
	}
	unitMap := valueobject.NewUnitMapFromUnitMatrix(unitMatrix)

	return unitMap, nil
}

func (gr *GameRoom) GetUnitsWithCoordinates(coordinates []valueobject.Coordinate) ([]valueobject.Unit, error) {
	units := make([]valueobject.Unit, 0)
	for _, coord := range coordinates {
		unit := gr.game.GetUnit(coord)
		units = append(units, unit)
	}

	return units, nil
}

func (gr *GameRoom) UpdateUnit(coordinate valueobject.Coordinate, unit valueobject.Unit) {
	gr.game.SetUnit(coordinate, unit)
}

func (gr *GameRoom) ReviveUnits(coordinates []valueobject.Coordinate) ([]valueobject.Coordinate, []valueobject.Unit, error) {
	if !gr.GetUnitMapSize().IncludesAllCoordinates(coordinates) {
		return nil, nil, ErrSomeCoordinatesNotIncludedInMap
	}
	updatedUnits := make([]valueobject.Unit, 0)
	for _, coordinate := range coordinates {
		unit := gr.game.GetUnit(coordinate)
		newUnit := unit.SetAlive(true)
		gr.game.SetUnit(coordinate, newUnit)
		updatedUnits = append(updatedUnits, newUnit)
	}
	return coordinates, updatedUnits, nil
}
