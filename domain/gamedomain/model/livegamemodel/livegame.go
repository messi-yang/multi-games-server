package livegamemodel

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/itemmodel"
)

var (
	ErrAreaExceedsUnitBlock            = errors.New("area should contain valid from and to coordinates and it should never exceed dimension")
	ErrSomeCoordinatesNotIncludedInMap = errors.New("some coordinates are not included in the unit map")
	ErrPlayerNotFound                  = errors.New("the play with the given id does not exist")
	ErrPlayerAlreadyExists             = errors.New("the play with the given id already exists")
)

type LiveGame struct {
	id          LiveGameId
	unitBlock   common.UnitBlock
	playerIds   map[common.PlayerId]bool
	zoomedAreas map[common.PlayerId]common.Area
}

func NewLiveGame(id LiveGameId, unitBlock common.UnitBlock) LiveGame {
	return LiveGame{
		id:          id,
		unitBlock:   unitBlock,
		playerIds:   make(map[common.PlayerId]bool),
		zoomedAreas: make(map[common.PlayerId]common.Area),
	}
}

func (liveGame *LiveGame) GetId() LiveGameId {
	return liveGame.id
}

func (liveGame *LiveGame) GetDimension() common.Dimension {
	return liveGame.unitBlock.GetDimension()
}

func (liveGame *LiveGame) GetUnitBlock() common.UnitBlock {
	return liveGame.unitBlock
}

func (liveGame *LiveGame) GetUnitBlockByArea(area common.Area) (common.UnitBlock, error) {
	if !liveGame.GetDimension().IncludesArea(area) {
		return common.UnitBlock{}, ErrAreaExceedsUnitBlock
	}
	offsetX := area.GetFrom().GetX()
	offsetY := area.GetFrom().GetY()
	areaWidth := area.GetWidth()
	areaHeight := area.GetHeight()
	unitMatrix := make([][]common.Unit, areaWidth)
	for x := 0; x < areaWidth; x += 1 {
		unitMatrix[x] = make([]common.Unit, areaHeight)
		for y := 0; y < areaHeight; y += 1 {
			coordinate, _ := common.NewCoordinate(x+offsetX, y+offsetY)
			unitMatrix[x][y] = liveGame.unitBlock.GetUnit(coordinate)
		}
	}
	unitBlock := common.NewUnitBlock(unitMatrix)

	return unitBlock, nil
}

func (liveGame *LiveGame) GetZoomedAreas() map[common.PlayerId]common.Area {
	return liveGame.zoomedAreas
}

func (liveGame *LiveGame) AddZoomedArea(playerId common.PlayerId, area common.Area) error {
	_, exists := liveGame.playerIds[playerId]
	if !exists {
		return ErrPlayerNotFound
	}
	liveGame.zoomedAreas[playerId] = area
	return nil
}

func (liveGame *LiveGame) RemoveZoomedArea(playerId common.PlayerId) {
	delete(liveGame.zoomedAreas, playerId)
}

func (liveGame *LiveGame) AddPlayer(playerId common.PlayerId) error {
	_, exists := liveGame.playerIds[playerId]
	if exists {
		return ErrPlayerAlreadyExists
	}

	liveGame.playerIds[playerId] = true

	return nil
}

func (liveGame *LiveGame) RemovePlayer(playerId common.PlayerId) {
	delete(liveGame.playerIds, playerId)
}

func (liveGame *LiveGame) BuildItem(coordinate common.Coordinate, itemId itemmodel.ItemId) error {
	if !liveGame.GetDimension().IncludesCoordinate(coordinate) {
		return ErrSomeCoordinatesNotIncludedInMap
	}

	unit := liveGame.unitBlock.GetUnit(coordinate)
	newUnit := unit.SetItemId(itemId)
	liveGame.unitBlock.SetUnit(coordinate, newUnit)

	return nil
}
