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
	ErrPlayerViewNotFound            = errors.New("ErrPlayerViewNotFound")
)

type LiveGame struct {
	id             LiveGameId
	map_           commonmodel.Map
	playerIds      map[playermodel.PlayerId]bool
	playerViews    map[playermodel.PlayerId]View
	observedRanges map[playermodel.PlayerId]commonmodel.Range
}

func NewLiveGame(id LiveGameId, map_ commonmodel.Map) LiveGame {
	return LiveGame{
		id:             id,
		map_:           map_,
		playerIds:      make(map[playermodel.PlayerId]bool),
		playerViews:    make(map[playermodel.PlayerId]View),
		observedRanges: make(map[playermodel.PlayerId]commonmodel.Range),
	}
}

func (liveGame *LiveGame) GetId() LiveGameId {
	return liveGame.id
}

func (liveGame *LiveGame) GetDimension() commonmodel.Dimension {
	return liveGame.map_.GetDimension()
}

func (liveGame *LiveGame) GetMap() commonmodel.Map {
	return liveGame.map_
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
		return liveGame.map_.GetUnit(location), nil
	})
	map_ := commonmodel.NewMap(unitMatrix)

	return map_, nil
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

func (liveGame *LiveGame) SetPlayerView(playerId playermodel.PlayerId, view View) {
	liveGame.playerViews[playerId] = view
}

func (liveGame *LiveGame) GetPlayerView(playerId playermodel.PlayerId) (View, error) {
	view, exists := liveGame.playerViews[playerId]
	if exists {
		return View{}, ErrPlayerViewNotFound
	}
	return view, nil
}

func (liveGame *LiveGame) AddPlayer(playerId playermodel.PlayerId) error {
	_, exists := liveGame.playerIds[playerId]
	if exists {
		return ErrPlayerAlreadyExists
	}

	liveGame.playerIds[playerId] = true

	originLocation, _ := commonmodel.NewLocation(0, 0)
	liveGame.SetPlayerView(playerId, NewView(originLocation))

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

	unit := liveGame.map_.GetUnit(location)
	newUnit := unit.SetItemId(itemId)
	liveGame.map_.ReplaceUnitAt(location, newUnit)

	return nil
}

func (liveGame *LiveGame) DestroyItem(location commonmodel.Location) error {
	if !liveGame.GetDimension().IncludesLocation(location) {
		return ErrSomeLocationsNotIncludedInMap
	}

	unit := liveGame.map_.GetUnit(location)
	itemId, _ := itemmodel.NewItemId(uuid.Nil.String())
	newUnit := unit.SetItemId(itemId)
	liveGame.map_.ReplaceUnitAt(location, newUnit)

	return nil
}
