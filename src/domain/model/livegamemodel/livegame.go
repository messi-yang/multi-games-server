package livegamemodel

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/playermodel"
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
	observedRanges map[playermodel.PlayerId]commonmodel.RangeVo
}

func NewLiveGame(id LiveGameId, unitMap commonmodel.UnitMap) LiveGame {
	return LiveGame{
		id:             id,
		unitMap:        unitMap,
		playerIds:      make(map[playermodel.PlayerId]bool),
		observedRanges: make(map[playermodel.PlayerId]commonmodel.RangeVo),
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

func (liveGame *LiveGame) GetUnitMapByRange(rangeVo commonmodel.RangeVo) (commonmodel.UnitMap, error) {
	if !liveGame.GetMapSize().IncludesRange(rangeVo) {
		return commonmodel.UnitMap{}, ErrRangeExceedsUnitMap
	}
	offsetX := rangeVo.GetFrom().GetX()
	offsetY := rangeVo.GetFrom().GetY()
	rangeVoWidth := rangeVo.GetWidth()
	rangeVoHeight := rangeVo.GetHeight()
	unitMatrix := make([][]commonmodel.Unit, rangeVoWidth)
	for x := 0; x < rangeVoWidth; x += 1 {
		unitMatrix[x] = make([]commonmodel.Unit, rangeVoHeight)
		for y := 0; y < rangeVoHeight; y += 1 {
			location, _ := commonmodel.NewLocation(x+offsetX, y+offsetY)
			unitMatrix[x][y] = liveGame.unitMap.GetUnit(location)
		}
	}
	unitMap := commonmodel.NewUnitMap(unitMatrix)

	return unitMap, nil
}

func (liveGame *LiveGame) GetObservedRanges() map[playermodel.PlayerId]commonmodel.RangeVo {
	return liveGame.observedRanges
}

func (liveGame *LiveGame) AddObservedRange(playerId playermodel.PlayerId, rangeVo commonmodel.RangeVo) error {
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
