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
	CameraChangedIntEventName         IntEventName = "CAMERA_CHANGED"
	ViewUpdatedIntEventName           IntEventName = "VIEW_UPDATED"
	ViewChangedIntEventName           IntEventName = "VIEW_CHANGED"
)

type GenericIntEvent struct {
	Name IntEventName `json:"name"`
}
type ChangeCameraRequestedintEvent struct {
	Name       IntEventName       `json:"name"`
	LiveGameId string             `json:"liveGameId"`
	PlayerId   string             `json:"playerId"`
	Camera     viewmodel.CameraVm `json:"camera"`
}

func NewChangeCameraRequestedintEvent(liveGameId string, playerId string, cameraVm viewmodel.CameraVm) ChangeCameraRequestedintEvent {
	return ChangeCameraRequestedintEvent{
		Name:       ChangeCameraRequestedIntEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		Camera:     cameraVm,
	}
}

type BuildItemRequestedintEvent struct {
	Name       IntEventName         `json:"name"`
	LiveGameId string               `json:"liveGameId"`
	Location   viewmodel.LocationVm `json:"location"`
	ItemId     string               `json:"itemId"`
}

func NewBuildItemRequestedintEvent(liveGameId string, locationVm viewmodel.LocationVm, itemId string) BuildItemRequestedintEvent {
	return BuildItemRequestedintEvent{
		Name:       BuildItemRequestedIntEventName,
		LiveGameId: liveGameId,
		Location:   locationVm,
		ItemId:     itemId,
	}
}

type JoinGameRequestedintEvent struct {
	Name       IntEventName `json:"name"`
	LiveGameId string       `json:"liveGameId"`
	PlayerId   string       `json:"playerId"`
}

func NewJoinGameRequestedintEvent(liveGameId string, playerId string) JoinGameRequestedintEvent {
	return JoinGameRequestedintEvent{
		Name:       JoinGameRequestedIntEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
	}
}

type DestroyItemRequestedintEvent struct {
	Name       IntEventName         `json:"name"`
	LiveGameId string               `json:"liveGameId"`
	Location   viewmodel.LocationVm `json:"location"`
}

func NewDestroyItemRequestedintEvent(liveGameId string, locationVm viewmodel.LocationVm) DestroyItemRequestedintEvent {
	return DestroyItemRequestedintEvent{
		Name:       DestroyItemRequestedIntEventName,
		LiveGameId: liveGameId,
		Location:   locationVm,
	}
}

type GameJoinedintEvent struct {
	Name       IntEventName       `json:"name"`
	LiveGameId string             `json:"liveGameId"`
	Player     viewmodel.PlayerVm `json:"playerId"`
	Camera     viewmodel.CameraVm `json:"camera"`
	MapSize    viewmodel.SizeVm   `json:"mapSize"`
	View       viewmodel.ViewVm   `json:"view"`
}

func NewGameJoinedintEvent(liveGameId string, playerVm viewmodel.PlayerVm, cameraVm viewmodel.CameraVm, mapSize viewmodel.SizeVm, viewVm viewmodel.ViewVm) GameJoinedintEvent {
	return GameJoinedintEvent{
		Name:       GameJoinedIntEventName,
		LiveGameId: liveGameId,
		Player:     playerVm,
		Camera:     cameraVm,
		MapSize:    mapSize,
		View:       viewVm,
	}
}

type CameraChangedintEvent struct {
	Name       IntEventName       `json:"name"`
	LiveGameId string             `json:"liveGameId"`
	PlayerId   string             `json:"playerId"`
	Camera     viewmodel.CameraVm `json:"camera"`
}

func NewCameraChangedintEvent(liveGameId string, playerId string, cameraVm viewmodel.CameraVm) CameraChangedintEvent {
	return CameraChangedintEvent{
		Name:       CameraChangedIntEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		Camera:     cameraVm,
	}
}

type ViewUpdatedintEvent struct {
	Name       IntEventName     `json:"name"`
	LiveGameId string           `json:"liveGameId"`
	PlayerId   string           `json:"playerId"`
	View       viewmodel.ViewVm `json:"view"`
}

func NewViewUpdatedintEvent(liveGameId string, playerId string, viewVm viewmodel.ViewVm) ViewUpdatedintEvent {
	return ViewUpdatedintEvent{
		Name:       ViewUpdatedIntEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		View:       viewVm,
	}
}

type ViewChangedintEvent struct {
	Name       IntEventName     `json:"name"`
	LiveGameId string           `json:"liveGameId"`
	PlayerId   string           `json:"playerId"`
	View       viewmodel.ViewVm `json:"view"`
}

func NewViewChangedintEvent(liveGameId string, playerId string, viewVm viewmodel.ViewVm) ViewChangedintEvent {
	return ViewChangedintEvent{
		Name:       ViewChangedIntEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		View:       viewVm,
	}
}

type LeaveGameRequestedintEvent struct {
	Name       IntEventName `json:"name"`
	LiveGameId string       `json:"liveGameId"`
	PlayerId   string       `json:"playerId"`
}

func NewLeaveGameRequestedintEvent(liveGameId string, playerId string) LeaveGameRequestedintEvent {
	return LeaveGameRequestedintEvent{
		Name:       LeaveGameRequestedIntEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
	}
}
