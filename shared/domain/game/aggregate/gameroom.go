package aggregate

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/entity"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/google/uuid"
)

var (
	ErrAreaExceedsUnitMap              = errors.New("area should contain valid from and to coordinates and it should never exceed map size")
	ErrSomeCoordinatesNotIncludedInMap = errors.New("some coordinates are not included in the unit map")
	ErrPlayerNotFound                  = errors.New("the play with the given id does not exist")
	ErrPlayerAlreadyExists             = errors.New("the play with the given id already exists")
)

type GameRoom struct {
	game        *entity.Game
	playerIds   map[uuid.UUID]bool
	zoomedAreas map[uuid.UUID]valueobject.Area
}

func NewGameRoom(game entity.Game) GameRoom {
	return GameRoom{
		game:        &game,
		playerIds:   make(map[uuid.UUID]bool),
		zoomedAreas: make(map[uuid.UUID]valueobject.Area),
	}
}

func (gr *GameRoom) GetId() uuid.UUID {
	return gr.game.GetId()
}

func (gr *GameRoom) GetMapSize() valueobject.MapSize {
	return gr.game.GetMapSize()
}

func (gr *GameRoom) GetUnitMap() valueobject.UnitMap {
	return gr.game.GetUnitMap()
}

func (gr *GameRoom) GetUnitMapByArea(area valueobject.Area) (valueobject.UnitMap, error) {
	if !gr.GetMapSize().IncludesArea(area) {
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
	unitMap := valueobject.NewUnitMap(unitMatrix)

	return unitMap, nil
}

func (gr *GameRoom) GetZoomedAreas() map[uuid.UUID]valueobject.Area {
	return gr.zoomedAreas
}

func (gr *GameRoom) AddZoomedArea(playerId uuid.UUID, area valueobject.Area) error {
	_, exists := gr.playerIds[playerId]
	if !exists {
		return ErrPlayerNotFound
	}
	gr.zoomedAreas[playerId] = area
	return nil
}

func (gr *GameRoom) RemoveZoomedArea(playerId uuid.UUID) {
	delete(gr.zoomedAreas, playerId)
}

func (gr *GameRoom) AddPlayer(playerId uuid.UUID) error {
	_, exists := gr.playerIds[playerId]
	if exists {
		return ErrPlayerAlreadyExists
	}

	gr.playerIds[playerId] = true

	return nil
}

func (gr *GameRoom) RemovePlayer(playerId uuid.UUID) {
	delete(gr.playerIds, playerId)
}

func (gr *GameRoom) ReviveUnits(coordinates []valueobject.Coordinate) error {
	if !gr.GetMapSize().IncludesAllCoordinates(coordinates) {
		return ErrSomeCoordinatesNotIncludedInMap
	}

	for _, coordinate := range coordinates {
		unit := gr.game.GetUnit(coordinate)
		newUnit := unit.SetAlive(true)
		gr.game.SetUnit(coordinate, newUnit)
	}

	return nil
}
