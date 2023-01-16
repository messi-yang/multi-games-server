package livegamemodel

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/playermodel"
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

type LiveGameAgr struct {
	id            LiveGameIdVo
	map_          MapVo
	playerIds     map[playermodel.PlayerIdVo]bool
	playerCameras map[playermodel.PlayerIdVo]CameraVo
}

func NewLiveGameAgr(id LiveGameIdVo, map_ MapVo) LiveGameAgr {
	return LiveGameAgr{
		id:            id,
		map_:          map_,
		playerIds:     make(map[playermodel.PlayerIdVo]bool),
		playerCameras: make(map[playermodel.PlayerIdVo]CameraVo),
	}
}

func (liveGame *LiveGameAgr) removePlayerCamera(playerId playermodel.PlayerIdVo) {
	delete(liveGame.playerCameras, playerId)
}

func (liveGame *LiveGameAgr) GetId() LiveGameIdVo {
	return liveGame.id
}

func (liveGame *LiveGameAgr) GetMapSize() commonmodel.SizeVo {
	return liveGame.map_.GetSize()
}

func (liveGame *LiveGameAgr) GetPlayerIds() []playermodel.PlayerIdVo {
	return lo.Keys(liveGame.playerCameras)
}

func (liveGame *LiveGameAgr) GetPlayerView(playerId playermodel.PlayerIdVo) (ViewVo, error) {
	camera, exists := liveGame.playerCameras[playerId]
	if !exists {
		return ViewVo{}, ErrPlayerCameraNotFound
	}

	view := liveGame.map_.GetViewWithCamera(camera)

	return view, nil
}

func (liveGame *LiveGameAgr) CanPlayerSeeAnyLocations(playerId playermodel.PlayerIdVo, locations []commonmodel.LocationVo) bool {
	camera, exists := liveGame.playerCameras[playerId]
	if !exists {
		return false
	}

	bound := camera.GetViewBoundInMap(liveGame.map_.GetSize())
	return bound.CoverAnyLocations(locations)
}

func (liveGame *LiveGameAgr) ChangePlayerCamera(playerId playermodel.PlayerIdVo, camera CameraVo) error {
	_, exists := liveGame.playerIds[playerId]
	if !exists {
		return ErrPlayerNotFound
	}

	liveGame.playerCameras[playerId] = camera
	return nil
}

func (liveGame *LiveGameAgr) GetPlayerCamera(playerId playermodel.PlayerIdVo) (CameraVo, error) {
	camera, exists := liveGame.playerCameras[playerId]
	if !exists {
		return CameraVo{}, ErrPlayerCameraNotFound
	}
	return camera, nil
}

func (liveGame *LiveGameAgr) AddPlayer(playerId playermodel.PlayerIdVo) error {
	_, exists := liveGame.playerIds[playerId]
	if exists {
		return ErrPlayerAlreadyExists
	}

	liveGame.playerIds[playerId] = true

	originLocation, _ := commonmodel.NewLocationVo(0, 0)
	liveGame.ChangePlayerCamera(playerId, NewCameraVo(originLocation))

	return nil
}

func (liveGame *LiveGameAgr) RemovePlayer(playerId playermodel.PlayerIdVo) {
	liveGame.removePlayerCamera(playerId)
	delete(liveGame.playerIds, playerId)
}

func (liveGame *LiveGameAgr) BuildItem(location commonmodel.LocationVo, itemId itemmodel.ItemIdVo) error {
	if !liveGame.GetMapSize().CoverLocation(location) {
		return ErrSomeLocationsNotIncludedInMap
	}

	unit := liveGame.map_.GetUnit(location)
	newUnit := unit.SetItemId(itemId)
	liveGame.map_.ReplaceUnitAt(location, newUnit)

	return nil
}

func (liveGame *LiveGameAgr) DestroyItem(location commonmodel.LocationVo) error {
	if !liveGame.GetMapSize().CoverLocation(location) {
		return ErrSomeLocationsNotIncludedInMap
	}

	unit := liveGame.map_.GetUnit(location)
	itemId, _ := itemmodel.NewItemIdVo(uuid.Nil.String())
	newUnit := unit.SetItemId(itemId)
	liveGame.map_.ReplaceUnitAt(location, newUnit)

	return nil
}
