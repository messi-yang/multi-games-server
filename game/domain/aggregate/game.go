package aggregate

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/game/domain/entity"
	"github.com/dum-dum-genius/game-of-liberty-computer/game/domain/valueobject"
	"github.com/google/uuid"
)

var (
	ErrAreaExceedsUnitBlock            = errors.New("area should contain valid from and to coordinates and it should never exceed dimension")
	ErrSomeCoordinatesNotIncludedInMap = errors.New("some coordinates are not included in the unit map")
	ErrPlayerNotFound                  = errors.New("the play with the given id does not exist")
	ErrPlayerAlreadyExists             = errors.New("the play with the given id already exists")
)

type Game struct {
	sandbox     entity.Sandbox
	playerIds   map[valueobject.PlayerId]bool
	zoomedAreas map[valueobject.PlayerId]valueobject.Area
}

func NewGame(sandbox entity.Sandbox) Game {
	return Game{
		sandbox:     sandbox,
		playerIds:   make(map[valueobject.PlayerId]bool),
		zoomedAreas: make(map[valueobject.PlayerId]valueobject.Area),
	}
}

func (gr *Game) GetId() uuid.UUID {
	return gr.sandbox.GetId()
}

func (gr *Game) GetDimension() valueobject.Dimension {
	return gr.sandbox.GetDimension()
}

func (gr *Game) GetUnitBlock() valueobject.UnitBlock {
	return gr.sandbox.GetUnitBlock()
}

func (gr *Game) GetUnitBlockByArea(area valueobject.Area) (valueobject.UnitBlock, error) {
	if !gr.GetDimension().IncludesArea(area) {
		return valueobject.UnitBlock{}, ErrAreaExceedsUnitBlock
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
			unitMatrix[x][y] = gr.sandbox.GetUnit(coordinate)
		}
	}
	unitBlock := valueobject.NewUnitBlock(unitMatrix)

	return unitBlock, nil
}

func (gr *Game) GetZoomedAreas() map[valueobject.PlayerId]valueobject.Area {
	return gr.zoomedAreas
}

func (gr *Game) AddZoomedArea(playerId valueobject.PlayerId, area valueobject.Area) error {
	_, exists := gr.playerIds[playerId]
	if !exists {
		return ErrPlayerNotFound
	}
	gr.zoomedAreas[playerId] = area
	return nil
}

func (gr *Game) RemoveZoomedArea(playerId valueobject.PlayerId) {
	delete(gr.zoomedAreas, playerId)
}

func (gr *Game) AddPlayer(playerId valueobject.PlayerId) error {
	_, exists := gr.playerIds[playerId]
	if exists {
		return ErrPlayerAlreadyExists
	}

	gr.playerIds[playerId] = true

	return nil
}

func (gr *Game) RemovePlayer(playerId valueobject.PlayerId) {
	delete(gr.playerIds, playerId)
}

func (gr *Game) ReviveUnits(coordinates []valueobject.Coordinate) error {
	if !gr.GetDimension().IncludesAllCoordinates(coordinates) {
		return ErrSomeCoordinatesNotIncludedInMap
	}

	for _, coordinate := range coordinates {
		unit := gr.sandbox.GetUnit(coordinate)
		newUnit := unit.SetAlive(true)
		gr.sandbox.SetUnit(coordinate, newUnit)
	}

	return nil
}
