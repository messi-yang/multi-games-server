package livegamemodel

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/playermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/tool"
	"github.com/google/uuid"
)

var (
	ErrRangeExceedsUnitMap           = errors.New("map range should contain valid from and to locations and it should never exceed mapSize")
	ErrSomeLocationsNotIncludedInMap = errors.New("some locations are not included in the unit map")
	ErrPlayerNotFound                = errors.New("the play with the given id does not exist")
	ErrPlayerAlreadyExists           = errors.New("the play with the given id already exists")
)

type LiveGame struct {
	id             LiveGameId
	unitMap        commonmodel.UnitMap
	playerIds      map[playermodel.PlayerId]bool
	observedRanges map[playermodel.PlayerId]commonmodel.Range
}

func NewLiveGame(id LiveGameId, unitMap commonmodel.UnitMap) LiveGame {
	return LiveGame{
		id:             id,
		unitMap:        unitMap,
		playerIds:      make(map[playermodel.PlayerId]bool),
		observedRanges: make(map[playermodel.PlayerId]commonmodel.Range),
	}
}

func (liveGame *LiveGame) GetId() LiveGameId {
	return liveGame.id
}

func (liveGame *LiveGame) GetMapSize() commonmodel.MapSize {
	return liveGame.unitMap.GetMapSize()
}

func (liveGame *LiveGame) GetUnitMap() commonmodel.UnitMap {
	return liveGame.unitMap
}

func (liveGame *LiveGame) GetUnitMapByRange(rangeVo commonmodel.Range) (commonmodel.UnitMap, error) {
	if !liveGame.GetMapSize().IncludesRange(rangeVo) {
		return commonmodel.UnitMap{}, ErrRangeExceedsUnitMap
	}
	offsetX := rangeVo.GetFrom().GetX()
	offsetY := rangeVo.GetFrom().GetY()
	rangeVoWidth := rangeVo.GetWidth()
	rangeVoHeight := rangeVo.GetHeight()
	unitMatrix, _ := tool.RangeMatrix(rangeVoWidth, rangeVoHeight, func(x int, y int) (commonmodel.Unit, error) {
		location, _ := commonmodel.NewLocation(x+offsetX, y+offsetY)
		return liveGame.unitMap.GetUnit(location), nil
	})
	unitMap := commonmodel.NewUnitMap(unitMatrix)

	return unitMap, nil
}

func (liveGame *LiveGame) GetObservedRanges() map[playermodel.PlayerId]commonmodel.Range {
	return liveGame.observedRanges
}

func (liveGame *LiveGame) AddObservedRange(playerId playermodel.PlayerId, rangeVo commonmodel.Range) error {
	_, exists := liveGame.playerIds[playerId]
	if !exists {
		return ErrPlayerNotFound
	}
	liveGame.observedRanges[playerId] = rangeVo
	return nil
}

func (liveGame *LiveGame) RemoveObservedRange(playerId playermodel.PlayerId) {
	delete(liveGame.observedRanges, playerId)
}

func (liveGame *LiveGame) AddPlayer(playerId playermodel.PlayerId) error {
	_, exists := liveGame.playerIds[playerId]
	if exists {
		return ErrPlayerAlreadyExists
	}

	liveGame.playerIds[playerId] = true

	return nil
}

func (liveGame *LiveGame) RemovePlayer(playerId playermodel.PlayerId) {
	delete(liveGame.playerIds, playerId)
}

func (liveGame *LiveGame) BuildItem(location commonmodel.Location, itemId itemmodel.ItemId) error {
	if !liveGame.GetMapSize().IncludesLocation(location) {
		return ErrSomeLocationsNotIncludedInMap
	}

	unit := liveGame.unitMap.GetUnit(location)
	newUnit := unit.SetItemId(itemId)
	liveGame.unitMap.ReplaceUnitAt(location, newUnit)

	return nil
}

func (liveGame *LiveGame) DestroyItem(location commonmodel.Location) error {
	if !liveGame.GetMapSize().IncludesLocation(location) {
		return ErrSomeLocationsNotIncludedInMap
	}

	unit := liveGame.unitMap.GetUnit(location)
	itemId, _ := itemmodel.NewItemId(uuid.Nil.String())
	newUnit := unit.SetItemId(itemId)
	liveGame.unitMap.ReplaceUnitAt(location, newUnit)

	return nil
}
