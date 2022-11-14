package livegamemodel

import (
	"errors"

	commonValueObject "github.com/dum-dum-genius/game-of-liberty-computer/common/domain/valueobject"
)

var (
	ErrAreaExceedsUnitBlock            = errors.New("area should contain valid from and to coordinates and it should never exceed dimension")
	ErrSomeCoordinatesNotIncludedInMap = errors.New("some coordinates are not included in the unit map")
	ErrPlayerNotFound                  = errors.New("the play with the given id does not exist")
	ErrPlayerAlreadyExists             = errors.New("the play with the given id already exists")
)

type LiveGame struct {
	id          LiveGameId
	unitBlock   commonValueObject.UnitBlock
	playerIds   map[commonValueObject.PlayerId]bool
	zoomedAreas map[commonValueObject.PlayerId]commonValueObject.Area
}

func NewLiveGame(id LiveGameId, unitBlock commonValueObject.UnitBlock) LiveGame {
	return LiveGame{
		id:          id,
		unitBlock:   unitBlock,
		playerIds:   make(map[commonValueObject.PlayerId]bool),
		zoomedAreas: make(map[commonValueObject.PlayerId]commonValueObject.Area),
	}
}

func (liveGame *LiveGame) GetId() LiveGameId {
	return liveGame.id
}

func (liveGame *LiveGame) GetDimension() commonValueObject.Dimension {
	return liveGame.unitBlock.GetDimension()
}

func (liveGame *LiveGame) GetUnitBlock() commonValueObject.UnitBlock {
	return liveGame.unitBlock
}

func (liveGame *LiveGame) GetUnitBlockByArea(area commonValueObject.Area) (commonValueObject.UnitBlock, error) {
	if !liveGame.GetDimension().IncludesArea(area) {
		return commonValueObject.UnitBlock{}, ErrAreaExceedsUnitBlock
	}
	offsetX := area.GetFrom().GetX()
	offsetY := area.GetFrom().GetY()
	areaWidth := area.GetWidth()
	areaHeight := area.GetHeight()
	unitMatrix := make([][]commonValueObject.Unit, areaWidth)
	for x := 0; x < areaWidth; x += 1 {
		unitMatrix[x] = make([]commonValueObject.Unit, areaHeight)
		for y := 0; y < areaHeight; y += 1 {
			coordinate, _ := commonValueObject.NewCoordinate(x+offsetX, y+offsetY)
			unitMatrix[x][y] = liveGame.unitBlock.GetUnit(coordinate)
		}
	}
	unitBlock := commonValueObject.NewUnitBlock(unitMatrix)

	return unitBlock, nil
}

func (liveGame *LiveGame) GetZoomedAreas() map[commonValueObject.PlayerId]commonValueObject.Area {
	return liveGame.zoomedAreas
}

func (liveGame *LiveGame) AddZoomedArea(playerId commonValueObject.PlayerId, area commonValueObject.Area) error {
	_, exists := liveGame.playerIds[playerId]
	if !exists {
		return ErrPlayerNotFound
	}
	liveGame.zoomedAreas[playerId] = area
	return nil
}

func (liveGame *LiveGame) RemoveZoomedArea(playerId commonValueObject.PlayerId) {
	delete(liveGame.zoomedAreas, playerId)
}

func (liveGame *LiveGame) AddPlayer(playerId commonValueObject.PlayerId) error {
	_, exists := liveGame.playerIds[playerId]
	if exists {
		return ErrPlayerAlreadyExists
	}

	liveGame.playerIds[playerId] = true

	return nil
}

func (liveGame *LiveGame) RemovePlayer(playerId commonValueObject.PlayerId) {
	delete(liveGame.playerIds, playerId)
}

func (liveGame *LiveGame) ReviveUnits(coordinates []commonValueObject.Coordinate) error {
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
