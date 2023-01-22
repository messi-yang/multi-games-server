package intevent

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
)

type IntEventName string

const (
	ChangeCameraRequestedIntEventName IntEventName = "CHANGE_CAMERA_REQUESTED"
	JoinGameRequestedIntEventName     IntEventName = "JOIN_GAME_REQUESTED"
	BuildItemRequestedIntEventName    IntEventName = "BUILD_ITEM_REQUESTED"
	DestroyItemRequestedIntEventName  IntEventName = "DESTROY_ITEM_REQUESTED"
	LeaveGameRequestedIntEventName    IntEventName = "LEAVE_GAME_REQUESTED"
	GameJoinedIntEventName            IntEventName = "GAME_JOINED"
	PlayerUpdatedIntEventName         IntEventName = "PLAYER_UPDATED"
	PlayersUpdatedIntEventName        IntEventName = "PLAYERS_UPDATED"
	ViewUpdatedIntEventName           IntEventName = "VIEW_UPDATED"
)

type GenericIntEvent struct {
	Name IntEventName `json:"name"`
}
type ChangeCameraRequestedIntEvent struct {
	Name       IntEventName       `json:"name"`
	LiveGameId string             `json:"liveGameId"`
	PlayerId   string             `json:"playerId"`
	Camera     viewmodel.CameraVm `json:"camera"`
}

func NewChangeCameraRequestedIntEvent(liveGameId string, playerId string, cameraVm viewmodel.CameraVm) ChangeCameraRequestedIntEvent {
	return ChangeCameraRequestedIntEvent{
		Name:       ChangeCameraRequestedIntEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		Camera:     cameraVm,
	}
}

type BuildItemRequestedIntEvent struct {
	Name       IntEventName         `json:"name"`
	LiveGameId string               `json:"liveGameId"`
	Location   viewmodel.LocationVm `json:"location"`
	ItemId     string               `json:"itemId"`
}

func NewBuildItemRequestedIntEvent(liveGameId string, locationVm viewmodel.LocationVm, itemId string) BuildItemRequestedIntEvent {
	return BuildItemRequestedIntEvent{
		Name:       BuildItemRequestedIntEventName,
		LiveGameId: liveGameId,
		Location:   locationVm,
		ItemId:     itemId,
	}
}

type JoinGameRequestedIntEvent struct {
	Name       IntEventName `json:"name"`
	LiveGameId string       `json:"liveGameId"`
	PlayerId   string       `json:"playerId"`
}

func NewJoinGameRequestedIntEvent(liveGameId string, playerId string) JoinGameRequestedIntEvent {
	return JoinGameRequestedIntEvent{
		Name:       JoinGameRequestedIntEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
	}
}

type DestroyItemRequestedIntEvent struct {
	Name       IntEventName         `json:"name"`
	LiveGameId string               `json:"liveGameId"`
	Location   viewmodel.LocationVm `json:"location"`
}

func NewDestroyItemRequestedIntEvent(liveGameId string, locationVm viewmodel.LocationVm) DestroyItemRequestedIntEvent {
	return DestroyItemRequestedIntEvent{
		Name:       DestroyItemRequestedIntEventName,
		LiveGameId: liveGameId,
		Location:   locationVm,
	}
}

type GameJoinedIntEvent struct {
	Name       IntEventName       `json:"name"`
	LiveGameId string             `json:"liveGameId"`
	Player     viewmodel.PlayerVm `json:"playerId"`
	MapSize    viewmodel.SizeVm   `json:"mapSize"`
	View       viewmodel.ViewVm   `json:"view"`
}

func NewGameJoinedIntEvent(liveGameId string, playerVm viewmodel.PlayerVm, mapSize viewmodel.SizeVm, viewVm viewmodel.ViewVm) GameJoinedIntEvent {
	return GameJoinedIntEvent{
		Name:       GameJoinedIntEventName,
		LiveGameId: liveGameId,
		Player:     playerVm,
		MapSize:    mapSize,
		View:       viewVm,
	}
}

type PlayerUpdatedIntEvent struct {
	Name       IntEventName       `json:"name"`
	LiveGameId string             `json:"liveGameId"`
	Player     viewmodel.PlayerVm `json:"player"`
}

func NewPlayerUpdatedIntEvent(liveGameId string, playerVm viewmodel.PlayerVm) PlayerUpdatedIntEvent {
	return PlayerUpdatedIntEvent{
		Name:       PlayerUpdatedIntEventName,
		LiveGameId: liveGameId,
		Player:     playerVm,
	}
}

type PlayersUpdatedIntEvent struct {
	Name       IntEventName         `json:"name"`
	LiveGameId string               `json:"liveGameId"`
	Players    []viewmodel.PlayerVm `json:"players"`
}

func NewPlayersUpdatedIntEvent(liveGameId string, playerVms []viewmodel.PlayerVm) PlayersUpdatedIntEvent {
	return PlayersUpdatedIntEvent{
		Name:       PlayersUpdatedIntEventName,
		LiveGameId: liveGameId,
		Players:    playerVms,
	}
}

type ViewUpdatedIntEvent struct {
	Name       IntEventName     `json:"name"`
	LiveGameId string           `json:"liveGameId"`
	PlayerId   string           `json:"playerId"`
	View       viewmodel.ViewVm `json:"view"`
}

func NewViewUpdatedIntEvent(liveGameId string, playerId string, viewVm viewmodel.ViewVm) ViewUpdatedIntEvent {
	return ViewUpdatedIntEvent{
		Name:       ViewUpdatedIntEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		View:       viewVm,
	}
}

type LeaveGameRequestedIntEvent struct {
	Name       IntEventName `json:"name"`
	LiveGameId string       `json:"liveGameId"`
	PlayerId   string       `json:"playerId"`
}

func NewLeaveGameRequestedIntEvent(liveGameId string, playerId string) LeaveGameRequestedIntEvent {
	return LeaveGameRequestedIntEvent{
		Name:       LeaveGameRequestedIntEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
	}
}
