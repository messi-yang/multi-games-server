package livegamemodel

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
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

type playerStatus struct {
	camera CameraVo
}

type LiveGameAgg struct {
	id             LiveGameIdVo
	map_           MapVo
	players        map[PlayerIdVo]PlayerEntity
	playerStatuses map[PlayerIdVo]playerStatus
}

func NewLiveGameAgg(id LiveGameIdVo, map_ MapVo) LiveGameAgg {
	return LiveGameAgg{
		id:             id,
		map_:           map_,
		players:        make(map[PlayerIdVo]PlayerEntity),
		playerStatuses: make(map[PlayerIdVo]playerStatus),
	}
}

func (liveGame *LiveGameAgg) GetId() LiveGameIdVo {
	return liveGame.id
}

func (liveGame *LiveGameAgg) GetMapSize() commonmodel.SizeVo {
	return liveGame.map_.GetSize()
}

func (liveGame *LiveGameAgg) GetPlayerIds() []PlayerIdVo {
	return lo.Keys(liveGame.playerStatuses)
}

func (liveGame *LiveGameAgg) GetPlayerView(playerId PlayerIdVo) (ViewVo, error) {
	playerStatus, exists := liveGame.playerStatuses[playerId]
	if !exists {
		return ViewVo{}, ErrPlayerCameraNotFound
	}

	view := liveGame.map_.GetViewWithCamera(playerStatus.camera)

	return view, nil
}

func (liveGame *LiveGameAgg) CanPlayerSeeAnyLocations(playerId PlayerIdVo, locations []commonmodel.LocationVo) bool {
	playerStatus, exists := liveGame.playerStatuses[playerId]
	if !exists {
		return false
	}

	bound := playerStatus.camera.GetViewBoundInMap(liveGame.map_.GetSize())
	return bound.CoverAnyLocations(locations)
}

func (liveGame *LiveGameAgg) ChangePlayerCamera(playerId PlayerIdVo, camera CameraVo) error {
	playerStatus, exists := liveGame.playerStatuses[playerId]
	if !exists {
		return ErrPlayerNotFound
	}

	playerStatus.camera = camera
	liveGame.playerStatuses[playerId] = playerStatus
	return nil
}

func (liveGame *LiveGameAgg) GetPlayerCamera(playerId PlayerIdVo) (CameraVo, error) {
	playerStatus, exists := liveGame.playerStatuses[playerId]
	if !exists {
		return CameraVo{}, ErrPlayerCameraNotFound
	}
	return playerStatus.camera, nil
}

func (liveGame *LiveGameAgg) AddPlayer(player PlayerEntity) error {
	_, exists := liveGame.players[player.GetId()]
	if exists {
		return ErrPlayerAlreadyExists
	}
	playerId := player.GetId()

	liveGame.players[playerId] = player

	originLocation, _ := commonmodel.NewLocationVo(0, 0)
	liveGame.playerStatuses[playerId] = playerStatus{
		camera: NewCameraVo(originLocation),
	}

	return nil
}

func (liveGame *LiveGameAgg) GetPlayer(playerId PlayerIdVo) (PlayerEntity, error) {
	player, exists := liveGame.players[playerId]
	if !exists {
		return PlayerEntity{}, ErrPlayerNotFound
	}
	return player, nil
}

func (liveGame *LiveGameAgg) RemovePlayer(playerId PlayerIdVo) {
	delete(liveGame.players, playerId)
	delete(liveGame.playerStatuses, playerId)
}

func (liveGame *LiveGameAgg) BuildItem(location commonmodel.LocationVo, itemId itemmodel.ItemIdVo) error {
	if !liveGame.GetMapSize().CoverLocation(location) {
		return ErrSomeLocationsNotIncludedInMap
	}

	unit := liveGame.map_.GetUnit(location)
	newUnit := unit.SetItemId(itemId)
	liveGame.map_.UpdateUnit(location, newUnit)

	return nil
}

func (liveGame *LiveGameAgg) DestroyItem(location commonmodel.LocationVo) error {
	if !liveGame.GetMapSize().CoverLocation(location) {
		return ErrSomeLocationsNotIncludedInMap
	}

	unit := liveGame.map_.GetUnit(location)
	itemId, _ := itemmodel.NewItemIdVo(uuid.Nil.String())
	newUnit := unit.SetItemId(itemId)
	liveGame.map_.UpdateUnit(location, newUnit)

	return nil
}
