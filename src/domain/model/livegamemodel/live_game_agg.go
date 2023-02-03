package livegamemodel

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

var (
	ErrSomeLocationsNotIncludedInMap = errors.New("some locations are not included in the unit map")
	ErrPlayerNotFound                = errors.New("the play with the given id does not exist")
	ErrPlayerAlreadyExists           = errors.New("the play with the given id already exists")
)

type LiveGameAgg struct {
	id      LiveGameIdVo
	map_    MapVo
	players map[PlayerIdVo]PlayerEntity
}

func NewLiveGameAgg(id LiveGameIdVo, map_ MapVo) LiveGameAgg {
	return LiveGameAgg{
		id:      id,
		map_:    map_,
		players: make(map[PlayerIdVo]PlayerEntity),
	}
}

func (liveGame *LiveGameAgg) GetId() LiveGameIdVo {
	return liveGame.id
}

func (liveGame *LiveGameAgg) GetMapSize() commonmodel.SizeVo {
	return liveGame.map_.GetSize()
}

func (liveGame *LiveGameAgg) GetUnit(location commonmodel.LocationVo) commonmodel.UnitVo {
	return liveGame.map_.GetUnit(location)
}

func (liveGame *LiveGameAgg) GetPlayerIds() []PlayerIdVo {
	return lo.Keys(liveGame.players)
}

func (liveGame *LiveGameAgg) GetPlayerIdsExcept(playerId PlayerIdVo) []PlayerIdVo {
	playerIds := lo.Keys(liveGame.players)
	return lo.Filter(playerIds, func(pId PlayerIdVo, _ int) bool {
		return !pId.isEqual(playerId)
	})
}

func (liveGame *LiveGameAgg) GetPlayerView(playerId PlayerIdVo) (ViewVo, error) {
	player, exists := liveGame.players[playerId]
	if !exists {
		return ViewVo{}, ErrPlayerNotFound
	}

	view := liveGame.map_.GetViewWithCamera(player.camera)

	return view, nil
}

func (liveGame *LiveGameAgg) CanPlayerSeeAnyLocations(playerId PlayerIdVo, locations []commonmodel.LocationVo) bool {
	player, exists := liveGame.players[playerId]
	if !exists {
		return false
	}

	bound := player.camera.GetViewBoundInMap(liveGame.map_.GetSize())
	return bound.CoverAnyLocations(locations)
}

func (liveGame *LiveGameAgg) ChangePlayerCamera(playerId PlayerIdVo, camera CameraVo) error {
	player, exists := liveGame.players[playerId]
	if !exists {
		return ErrPlayerNotFound
	}

	player.camera = camera
	liveGame.players[playerId] = player
	return nil
}

func (liveGame *LiveGameAgg) AddPlayer(playerId PlayerIdVo) error {
	_, exists := liveGame.players[playerId]
	if exists {
		return ErrPlayerAlreadyExists
	}

	cameraCenter := commonmodel.NewLocationVo(0, 0)
	camera := NewCameraVo(cameraCenter)

	playerLocation := commonmodel.NewLocationVo(0, 0)
	newPlayer := NewPlayerEntity(playerId, "Hello World", camera, playerLocation)

	liveGame.players[playerId] = newPlayer

	return nil
}

func (liveGame *LiveGameAgg) UpdatePlayer(player PlayerEntity) error {
	_, exists := liveGame.players[player.id]
	if !exists {
		return ErrPlayerNotFound
	}
	liveGame.players[player.id] = player
	return nil
}

func (liveGame *LiveGameAgg) GetPlayers() []PlayerEntity {
	return lo.Values(liveGame.players)
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
}

func (liveGame *LiveGameAgg) BuildItem(location commonmodel.LocationVo, itemId itemmodel.ItemIdVo) error {
	if !liveGame.GetMapSize().CoversLocation(location) {
		return ErrSomeLocationsNotIncludedInMap
	}

	unit := liveGame.map_.GetUnit(location)
	newUnit := unit.SetItemId(itemId)
	liveGame.map_.UpdateUnit(location, newUnit)

	return nil
}

func (liveGame *LiveGameAgg) DestroyItem(location commonmodel.LocationVo) error {
	if !liveGame.GetMapSize().CoversLocation(location) {
		return ErrSomeLocationsNotIncludedInMap
	}

	unit := liveGame.map_.GetUnit(location)
	itemId, _ := itemmodel.NewItemIdVo(uuid.Nil.String())
	newUnit := unit.SetItemId(itemId)
	liveGame.map_.UpdateUnit(location, newUnit)

	return nil
}
