package livegamemodel

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/playermodel"
	"github.com/google/uuid"
)

var (
	ErrExtentExceedsUnitMap          = errors.New("map range should contain valid from and to locations and it should never exceed mapSize")
	ErrSomeLocationsNotIncludedInMap = errors.New("some locations are not included in the unit map")
	ErrPlayerNotFound                = errors.New("the play with the given id does not exist")
	ErrPlayerAlreadyExists           = errors.New("the play with the given id already exists")
)

type LiveGame struct {
	id              LiveGameId
	unitMap         commonmodel.UnitMap
	playerIds       map[playermodel.PlayerId]bool
	observedExtents map[playermodel.PlayerId]commonmodel.Extent
}

func NewLiveGame(id LiveGameId, unitMap commonmodel.UnitMap) LiveGame {
	return LiveGame{
		id:              id,
		unitMap:         unitMap,
		playerIds:       make(map[playermodel.PlayerId]bool),
		observedExtents: make(map[playermodel.PlayerId]commonmodel.Extent),
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

func (liveGame *LiveGame) GetUnitMapByExtent(extent commonmodel.Extent) (commonmodel.UnitMap, error) {
	if !liveGame.GetMapSize().IncludesExtent(extent) {
		return commonmodel.UnitMap{}, ErrExtentExceedsUnitMap
	}
	offsetX := extent.GetFrom().GetX()
	offsetY := extent.GetFrom().GetY()
	extentWidth := extent.GetWidth()
	extentHeight := extent.GetHeight()
	unitMatrix := make([][]commonmodel.Unit, extentWidth)
	for x := 0; x < extentWidth; x += 1 {
		unitMatrix[x] = make([]commonmodel.Unit, extentHeight)
		for y := 0; y < extentHeight; y += 1 {
			location, _ := commonmodel.NewLocation(x+offsetX, y+offsetY)
			unitMatrix[x][y] = liveGame.unitMap.GetUnit(location)
		}
	}
	unitMap := commonmodel.NewUnitMap(unitMatrix)

	return unitMap, nil
}

func (liveGame *LiveGame) GetObservedExtents() map[playermodel.PlayerId]commonmodel.Extent {
	return liveGame.observedExtents
}

func (liveGame *LiveGame) AddObservedExtent(playerId playermodel.PlayerId, extent commonmodel.Extent) error {
	_, exists := liveGame.playerIds[playerId]
	if !exists {
		return ErrPlayerNotFound
	}
	liveGame.observedExtents[playerId] = extent
	return nil
}

func (liveGame *LiveGame) RemoveObservedExtent(playerId playermodel.PlayerId) {
	delete(liveGame.observedExtents, playerId)
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
