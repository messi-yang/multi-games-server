package livegamemodel

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/playermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/tool"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

var (
	ErrBoundExceedsMap               = errors.New("map bound should contain valid from and to locations and it should never exceed size")
	ErrSomeLocationsNotIncludedInMap = errors.New("some locations are not included in the unit map")
	ErrPlayerNotFound                = errors.New("the play with the given id does not exist")
	ErrPlayerAlreadyExists           = errors.New("the play with the given id already exists")
	ErrPlayerCameraNotFound          = errors.New("ErrPlayerCameraNotFound")
)

type LiveGame struct {
	id            LiveGameId
	map_          Map
	playerIds     map[playermodel.PlayerId]bool
	playerCameras map[playermodel.PlayerId]Camera
}

func NewLiveGame(id LiveGameId, map_ Map) LiveGame {
	return LiveGame{
		id:            id,
		map_:          map_,
		playerIds:     make(map[playermodel.PlayerId]bool),
		playerCameras: make(map[playermodel.PlayerId]Camera),
	}
}

func (liveGame *LiveGame) getViewBoundOfCamera(camera Camera) Bound {
	fromX := camera.GetCenter().GetX() - 25
	toX := camera.GetCenter().GetX() + 25
	mapSize := liveGame.GetMapSize()
	mapWidth := mapSize.GetWidth()
	if fromX < 0 {
		toX -= fromX
		fromX = 0
	} else if toX > mapWidth-1 {
		fromX -= toX - mapWidth - 1
		toX = mapWidth - 1
	}

	fromY := camera.GetCenter().GetY() - 25
	toY := camera.GetCenter().GetY() + 25
	mapHeight := mapSize.GetHeight()
	if fromY < 0 {
		toY -= fromY
		fromY = 0
	} else if toY > mapHeight-1 {
		fromY -= toY - mapHeight - 1
		toY = mapHeight - 1
	}

	from, _ := commonmodel.NewLocation(fromX, fromY)
	to, _ := commonmodel.NewLocation(toX, toY)
	bound, _ := NewBound(from, to)

	return bound
}

func (liveGame *LiveGame) getViewOfCamera(camera Camera) View {
	bound := liveGame.getViewBoundOfCamera(camera)
	offsetX := bound.GetFrom().GetX()
	offsetY := bound.GetFrom().GetY()
	boundWidth := bound.GetWidth()
	boundHeight := bound.GetHeight()
	unitMatrix, _ := tool.RangeMatrix(boundWidth, boundHeight, func(x int, y int) (commonmodel.Unit, error) {
		location, _ := commonmodel.NewLocation(x+offsetX, y+offsetY)
		return liveGame.map_.GetUnit(location), nil
	})
	return NewView(NewMap(unitMatrix), bound)
}

func (liveGame *LiveGame) removePlayerCamera(playerId playermodel.PlayerId) {
	delete(liveGame.playerCameras, playerId)
}

func (liveGame *LiveGame) GetId() LiveGameId {
	return liveGame.id
}

func (liveGame *LiveGame) GetMapSize() commonmodel.Size {
	return liveGame.map_.GetSize()
}

func (liveGame *LiveGame) GetPlayerIds() []playermodel.PlayerId {
	return lo.Keys(liveGame.playerCameras)
}

func (liveGame *LiveGame) GetPlayerView(playerId playermodel.PlayerId) (View, error) {
	camera, exists := liveGame.playerCameras[playerId]
	if !exists {
		return View{}, ErrPlayerCameraNotFound
	}

	view := liveGame.getViewOfCamera(camera)

	return view, nil
}

func (liveGame *LiveGame) CanPlayerSeeAnyLocations(playerId playermodel.PlayerId, locations []commonmodel.Location) bool {
	camera, exists := liveGame.playerCameras[playerId]
	if !exists {
		return false
	}

	bound := liveGame.getViewBoundOfCamera(camera)
	return bound.CoverAnyLocations(locations)
}

func (liveGame *LiveGame) ChangePlayerCamera(playerId playermodel.PlayerId, camera Camera) error {
	_, exists := liveGame.playerIds[playerId]
	if !exists {
		return ErrPlayerNotFound
	}

	liveGame.playerCameras[playerId] = camera
	return nil
}

func (liveGame *LiveGame) GetPlayerCamera(playerId playermodel.PlayerId) (Camera, error) {
	camera, exists := liveGame.playerCameras[playerId]
	if !exists {
		return Camera{}, ErrPlayerCameraNotFound
	}
	return camera, nil
}

func (liveGame *LiveGame) AddPlayer(playerId playermodel.PlayerId) error {
	_, exists := liveGame.playerIds[playerId]
	if exists {
		return ErrPlayerAlreadyExists
	}

	liveGame.playerIds[playerId] = true

	originLocation, _ := commonmodel.NewLocation(0, 0)
	liveGame.ChangePlayerCamera(playerId, NewCamera(originLocation))

	return nil
}

func (liveGame *LiveGame) RemovePlayer(playerId playermodel.PlayerId) {
	liveGame.removePlayerCamera(playerId)
	delete(liveGame.playerIds, playerId)
}

func (liveGame *LiveGame) BuildItem(location commonmodel.Location, itemId itemmodel.ItemId) error {
	if !liveGame.GetMapSize().CoverLocation(location) {
		return ErrSomeLocationsNotIncludedInMap
	}

	unit := liveGame.map_.GetUnit(location)
	newUnit := unit.SetItemId(itemId)
	liveGame.map_.ReplaceUnitAt(location, newUnit)

	return nil
}

func (liveGame *LiveGame) DestroyItem(location commonmodel.Location) error {
	if !liveGame.GetMapSize().CoverLocation(location) {
		return ErrSomeLocationsNotIncludedInMap
	}

	unit := liveGame.map_.GetUnit(location)
	itemId, _ := itemmodel.NewItemId(uuid.Nil.String())
	newUnit := unit.SetItemId(itemId)
	liveGame.map_.ReplaceUnitAt(location, newUnit)

	return nil
}
