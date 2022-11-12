package aggregate

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/domain/valueobject"
)

var (
	ErrAreaExceedsUnitBlock            = errors.New("area should contain valid from and to coordinates and it should never exceed dimension")
	ErrSomeCoordinatesNotIncludedInMap = errors.New("some coordinates are not included in the unit map")
	ErrPlayerNotFound                  = errors.New("the play with the given id does not exist")
	ErrPlayerAlreadyExists             = errors.New("the play with the given id already exists")
)

type LiveGame struct {
	id          valueobject.GameId
	unitBlock   valueobject.UnitBlock
	playerIds   map[valueobject.PlayerId]bool
	zoomedAreas map[valueobject.PlayerId]valueobject.Area
}

func NewLiveGame(id valueobject.GameId, unitBlock valueobject.UnitBlock) LiveGame {
	return LiveGame{
		id:          id,
		unitBlock:   unitBlock,
		playerIds:   make(map[valueobject.PlayerId]bool),
		zoomedAreas: make(map[valueobject.PlayerId]valueobject.Area),
	}
}

func (liveGame *LiveGame) GetId() valueobject.GameId {
	return liveGame.id
}

func (liveGame *LiveGame) GetDimension() valueobject.Dimension {
	return liveGame.unitBlock.GetDimension()
}

func (liveGame *LiveGame) GetUnitBlock() valueobject.UnitBlock {
	return liveGame.unitBlock
}

func (liveGame *LiveGame) GetUnitBlockByArea(area valueobject.Area) (valueobject.UnitBlock, error) {
	if !liveGame.GetDimension().IncludesArea(area) {
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
			unitMatrix[x][y] = liveGame.unitBlock.GetUnit(coordinate)
		}
	}
	unitBlock := valueobject.NewUnitBlock(unitMatrix)

	return unitBlock, nil
}

func (liveGame *LiveGame) GetZoomedAreas() map[valueobject.PlayerId]valueobject.Area {
	return liveGame.zoomedAreas
}

func (liveGame *LiveGame) AddZoomedArea(playerId valueobject.PlayerId, area valueobject.Area) error {
	_, exists := liveGame.playerIds[playerId]
	if !exists {
		return ErrPlayerNotFound
	}
	liveGame.zoomedAreas[playerId] = area
	return nil
}

func (liveGame *LiveGame) RemoveZoomedArea(playerId valueobject.PlayerId) {
	delete(liveGame.zoomedAreas, playerId)
}

func (liveGame *LiveGame) AddPlayer(playerId valueobject.PlayerId) error {
	_, exists := liveGame.playerIds[playerId]
	if exists {
		return ErrPlayerAlreadyExists
	}

	liveGame.playerIds[playerId] = true

	return nil
}

func (liveGame *LiveGame) RemovePlayer(playerId valueobject.PlayerId) {
	delete(liveGame.playerIds, playerId)
}

func (liveGame *LiveGame) ReviveUnits(coordinates []valueobject.Coordinate) error {
	if !liveGame.GetDimension().IncludesAllCoordinates(coordinates) {
		return ErrSomeCoordinatesNotIncludedInMap
	}

	for _, coordinate := range coordinates {
		unit := liveGame.unitBlock.GetUnit(coordinate)
		newUnit := unit.SetAlive(true)
		liveGame.unitBlock.SetUnit(coordinate, newUnit)
	}

	return nil
}
