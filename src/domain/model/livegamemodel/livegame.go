package livegamemodel

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/playermodel"
	"github.com/google/uuid"
)

var (
	ErrMapRangeExceedsGameMap        = errors.New("map range should contain valid from and to locations and it should never exceed mapSize")
	ErrSomeLocationsNotIncludedInMap = errors.New("some locations are not included in the mapUnit map")
	ErrPlayerNotFound                = errors.New("the play with the given id does not exist")
	ErrPlayerAlreadyExists           = errors.New("the play with the given id already exists")
)

type LiveGame struct {
	id                LiveGameId
	gameMap           commonmodel.GameMap
	playerIds         map[playermodel.PlayerId]bool
	observedMapRanges map[playermodel.PlayerId]commonmodel.MapRange
}

func NewLiveGame(id LiveGameId, gameMap commonmodel.GameMap) LiveGame {
	return LiveGame{
		id:                id,
		gameMap:           gameMap,
		playerIds:         make(map[playermodel.PlayerId]bool),
		observedMapRanges: make(map[playermodel.PlayerId]commonmodel.MapRange),
	}
}

func (liveGame *LiveGame) GetId() LiveGameId {
	return liveGame.id
}

func (liveGame *LiveGame) GetMapSize() commonmodel.MapSize {
	return liveGame.gameMap.GetMapSize()
}

func (liveGame *LiveGame) GetGameMap() commonmodel.GameMap {
	return liveGame.gameMap
}

func (liveGame *LiveGame) GetGameMapByMapRange(mapRange commonmodel.MapRange) (commonmodel.GameMap, error) {
	if !liveGame.GetMapSize().IncludesMapRange(mapRange) {
		return commonmodel.GameMap{}, ErrMapRangeExceedsGameMap
	}
	offsetX := mapRange.GetFrom().GetX()
	offsetY := mapRange.GetFrom().GetY()
	mapRangeWidth := mapRange.GetWidth()
	mapRangeHeight := mapRange.GetHeight()
	mapUnitMatrix := make([][]commonmodel.MapUnit, mapRangeWidth)
	for x := 0; x < mapRangeWidth; x += 1 {
		mapUnitMatrix[x] = make([]commonmodel.MapUnit, mapRangeHeight)
		for y := 0; y < mapRangeHeight; y += 1 {
			location, _ := commonmodel.NewLocation(x+offsetX, y+offsetY)
			mapUnitMatrix[x][y] = liveGame.gameMap.GetMapUnit(location)
		}
	}
	gameMap := commonmodel.NewGameMap(mapUnitMatrix)

	return gameMap, nil
}

func (liveGame *LiveGame) GetObservedMapRanges() map[playermodel.PlayerId]commonmodel.MapRange {
	return liveGame.observedMapRanges
}

func (liveGame *LiveGame) AddObservedMapRange(playerId playermodel.PlayerId, mapRange commonmodel.MapRange) error {
	_, exists := liveGame.playerIds[playerId]
	if !exists {
		return ErrPlayerNotFound
	}
	liveGame.observedMapRanges[playerId] = mapRange
	return nil
}

func (liveGame *LiveGame) RemoveObservedMapRange(playerId playermodel.PlayerId) {
	delete(liveGame.observedMapRanges, playerId)
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

	mapUnit := liveGame.gameMap.GetMapUnit(location)
	newMapUnit := mapUnit.SetItemId(itemId)
	liveGame.gameMap.ReplaceMapUnitAt(location, newMapUnit)

	return nil
}

func (liveGame *LiveGame) DestroyItem(location commonmodel.Location) error {
	if !liveGame.GetMapSize().IncludesLocation(location) {
		return ErrSomeLocationsNotIncludedInMap
	}

	mapUnit := liveGame.gameMap.GetMapUnit(location)
	itemId, _ := itemmodel.NewItemId(uuid.Nil.String())
	newMapUnit := mapUnit.SetItemId(itemId)
	liveGame.gameMap.ReplaceMapUnitAt(location, newMapUnit)

	return nil
}
