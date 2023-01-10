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
	ErrRangeExceedsMap               = errors.New("map range should contain valid from and to locations and it should never exceed dimension")
	ErrSomeLocationsNotIncludedInMap = errors.New("some locations are not included in the unit map")
	ErrPlayerNotFound                = errors.New("the play with the given id does not exist")
	ErrPlayerAlreadyExists           = errors.New("the play with the given id already exists")
)

type LiveGame struct {
	id             LiveGameId
	mapVo          commonmodel.Map
	playerIds      map[playermodel.PlayerId]bool
	observedRanges map[playermodel.PlayerId]commonmodel.Range
}

func NewLiveGame(id LiveGameId, mapVo commonmodel.Map) LiveGame {
	return LiveGame{
		id:             id,
		mapVo:          mapVo,
		playerIds:      make(map[playermodel.PlayerId]bool),
		observedRanges: make(map[playermodel.PlayerId]commonmodel.Range),
	}
}

func (liveGame *LiveGame) GetId() LiveGameId {
	return liveGame.id
}

func (liveGame *LiveGame) GetDimension() commonmodel.Dimension {
	return liveGame.mapVo.GetDimension()
}

func (liveGame *LiveGame) GetMap() commonmodel.Map {
	return liveGame.mapVo
}

func (liveGame *LiveGame) GetMapByRange(rangeVo commonmodel.Range) (commonmodel.Map, error) {
	if !liveGame.GetDimension().IncludesRange(rangeVo) {
		return commonmodel.Map{}, ErrRangeExceedsMap
	}
	offsetX := rangeVo.GetFrom().GetX()
	offsetY := rangeVo.GetFrom().GetY()
	rangeVoWidth := rangeVo.GetWidth()
	rangeVoHeight := rangeVo.GetHeight()
	unitMatrix, _ := tool.RangeMatrix(rangeVoWidth, rangeVoHeight, func(x int, y int) (commonmodel.Unit, error) {
		location, _ := commonmodel.NewLocation(x+offsetX, y+offsetY)
		return liveGame.mapVo.GetUnit(location), nil
	})
	mapVo := commonmodel.NewMap(unitMatrix)

	return mapVo, nil
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

func (liveGame *LiveGame) removeObservedRange(playerId playermodel.PlayerId) {
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
	liveGame.removeObservedRange(playerId)
	delete(liveGame.playerIds, playerId)
}

func (liveGame *LiveGame) BuildItem(location commonmodel.Location, itemId itemmodel.ItemId) error {
	if !liveGame.GetDimension().IncludesLocation(location) {
		return ErrSomeLocationsNotIncludedInMap
	}

	unit := liveGame.mapVo.GetUnit(location)
	newUnit := unit.SetItemId(itemId)
	liveGame.mapVo.ReplaceUnitAt(location, newUnit)

	return nil
}

func (liveGame *LiveGame) DestroyItem(location commonmodel.Location) error {
	if !liveGame.GetDimension().IncludesLocation(location) {
		return ErrSomeLocationsNotIncludedInMap
	}

	unit := liveGame.mapVo.GetUnit(location)
	itemId, _ := itemmodel.NewItemId(uuid.Nil.String())
	newUnit := unit.SetItemId(itemId)
	liveGame.mapVo.ReplaceUnitAt(location, newUnit)

	return nil
}
