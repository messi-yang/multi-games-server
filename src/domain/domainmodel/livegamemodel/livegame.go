package livegamemodel

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/domainmodel/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/domainmodel/itemmodel"
	"github.com/google/uuid"
)

var (
	ErrAreaExceedsUnitBlock            = errors.New("area should contain valid from and to coordinates and it should never exceed dimension")
	ErrSomeCoordinatesNotIncludedInMap = errors.New("some coordinates are not included in the unit map")
	ErrPlayerNotFound                  = errors.New("the play with the given id does not exist")
	ErrPlayerAlreadyExists             = errors.New("the play with the given id already exists")
)

type LiveGame struct {
	id          LiveGameId
	unitBlock   commonmodel.UnitBlock
	playerIds   map[commonmodel.PlayerId]bool
	zoomedAreas map[commonmodel.PlayerId]commonmodel.Area
}

func NewLiveGame(id LiveGameId, unitBlock commonmodel.UnitBlock) LiveGame {
	return LiveGame{
		id:          id,
		unitBlock:   unitBlock,
		playerIds:   make(map[commonmodel.PlayerId]bool),
		zoomedAreas: make(map[commonmodel.PlayerId]commonmodel.Area),
	}
}

func (liveGame *LiveGame) GetId() LiveGameId {
	return liveGame.id
}

func (liveGame *LiveGame) GetDimension() commonmodel.Dimension {
	return liveGame.unitBlock.GetDimension()
}

func (liveGame *LiveGame) GetUnitBlock() commonmodel.UnitBlock {
	return liveGame.unitBlock
}

func (liveGame *LiveGame) GetUnitBlockByArea(area commonmodel.Area) (commonmodel.UnitBlock, error) {
	if !liveGame.GetDimension().IncludesArea(area) {
		return commonmodel.UnitBlock{}, ErrAreaExceedsUnitBlock
	}
	offsetX := area.GetFrom().GetX()
	offsetY := area.GetFrom().GetY()
	areaWidth := area.GetWidth()
	areaHeight := area.GetHeight()
	unitMatrix := make([][]commonmodel.Unit, areaWidth)
	for x := 0; x < areaWidth; x += 1 {
		unitMatrix[x] = make([]commonmodel.Unit, areaHeight)
		for y := 0; y < areaHeight; y += 1 {
			coordinate, _ := commonmodel.NewCoordinate(x+offsetX, y+offsetY)
			unitMatrix[x][y] = liveGame.unitBlock.GetUnit(coordinate)
		}
	}
	unitBlock := commonmodel.NewUnitBlock(unitMatrix)

	return unitBlock, nil
}

func (liveGame *LiveGame) GetZoomedAreas() map[commonmodel.PlayerId]commonmodel.Area {
	return liveGame.zoomedAreas
}

func (liveGame *LiveGame) AddZoomedArea(playerId commonmodel.PlayerId, area commonmodel.Area) error {
	_, exists := liveGame.playerIds[playerId]
	if !exists {
		return ErrPlayerNotFound
	}
	liveGame.zoomedAreas[playerId] = area
	return nil
}

func (liveGame *LiveGame) RemoveZoomedArea(playerId commonmodel.PlayerId) {
	delete(liveGame.zoomedAreas, playerId)
}

func (liveGame *LiveGame) AddPlayer(playerId commonmodel.PlayerId) error {
	_, exists := liveGame.playerIds[playerId]
	if exists {
		return ErrPlayerAlreadyExists
	}

	liveGame.playerIds[playerId] = true

	return nil
}

func (liveGame *LiveGame) RemovePlayer(playerId commonmodel.PlayerId) {
	delete(liveGame.playerIds, playerId)
}

func (liveGame *LiveGame) BuildItem(coordinate commonmodel.Coordinate, itemId itemmodel.ItemId) error {
	if !liveGame.GetDimension().IncludesCoordinate(coordinate) {
		return ErrSomeCoordinatesNotIncludedInMap
	}

	unit := liveGame.unitBlock.GetUnit(coordinate)
	newUnit := unit.SetItemId(itemId)
	liveGame.unitBlock.SetUnit(coordinate, newUnit)

	return nil
}

func (liveGame *LiveGame) DestroyItem(coordinate commonmodel.Coordinate) error {
	if !liveGame.GetDimension().IncludesCoordinate(coordinate) {
		return ErrSomeCoordinatesNotIncludedInMap
	}

	unit := liveGame.unitBlock.GetUnit(coordinate)
	itemId, _ := itemmodel.NewItemId(uuid.Nil.String())
	newUnit := unit.SetItemId(itemId)
	liveGame.unitBlock.SetUnit(coordinate, newUnit)

	return nil
}
