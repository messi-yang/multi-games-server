package intgrevent

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
)

type IntgrEventName string

const (
	ChangeCameraRequestedIntgrEventName IntgrEventName = "CHANGE_CAMERA_REQUESTED"
	JoinGameRequestedIntgrEventName     IntgrEventName = "JOIN_GAME_REQUESTED"
	BuildItemRequestedIntgrEventName    IntgrEventName = "BUILD_ITEM_REQUESTED"
	DestroyItemRequestedIntgrEventName  IntgrEventName = "DESTROY_ITEM_REQUESTED"
	LeaveGameRequestedIntgrEventName    IntgrEventName = "LEAVE_GAME_REQUESTED"
	GameJoinedIntgrEventName            IntgrEventName = "GAME_JOINED"
	CameraChangedIntgrEventName         IntgrEventName = "CAMERA_CHANGED"
	ViewUpdatedIntgrEventName           IntgrEventName = "VIEW_UPDATED"
	ViewChangedIntgrEventName           IntgrEventName = "VIEW_CHANGED"
)

type GenericIntgrEvent struct {
	Name IntgrEventName `json:"name"`
}
type ChangeCameraRequestedIntgrEvent struct {
	Name       IntgrEventName     `json:"name"`
	LiveGameId string             `json:"liveGameId"`
	PlayerId   string             `json:"playerId"`
	Camera     viewmodel.CameraVm `json:"camera"`
}

func NewChangeCameraRequestedIntgrEvent(liveGameId string, playerId string, cameraVm viewmodel.CameraVm) ChangeCameraRequestedIntgrEvent {
	return ChangeCameraRequestedIntgrEvent{
		Name:       ChangeCameraRequestedIntgrEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		Camera:     cameraVm,
	}
}

type BuildItemRequestedIntgrEvent struct {
	Name       IntgrEventName       `json:"name"`
	LiveGameId string               `json:"liveGameId"`
	Location   viewmodel.LocationVm `json:"location"`
	ItemId     string               `json:"itemId"`
}

func NewBuildItemRequestedIntgrEvent(liveGameId string, locationVm viewmodel.LocationVm, itemId string) BuildItemRequestedIntgrEvent {
	return BuildItemRequestedIntgrEvent{
		Name:       BuildItemRequestedIntgrEventName,
		LiveGameId: liveGameId,
		Location:   locationVm,
		ItemId:     itemId,
	}
}

type JoinGameRequestedIntgrEvent struct {
	Name       IntgrEventName `json:"name"`
	LiveGameId string         `json:"liveGameId"`
	PlayerId   string         `json:"playerId"`
}

func NewJoinGameRequestedIntgrEvent(liveGameId string, playerId string) JoinGameRequestedIntgrEvent {
	return JoinGameRequestedIntgrEvent{
		Name:       JoinGameRequestedIntgrEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
	}
}

type DestroyItemRequestedIntgrEvent struct {
	Name       IntgrEventName       `json:"name"`
	LiveGameId string               `json:"liveGameId"`
	Location   viewmodel.LocationVm `json:"location"`
}

func NewDestroyItemRequestedIntgrEvent(liveGameId string, locationVm viewmodel.LocationVm) DestroyItemRequestedIntgrEvent {
	return DestroyItemRequestedIntgrEvent{
		Name:       DestroyItemRequestedIntgrEventName,
		LiveGameId: liveGameId,
		Location:   locationVm,
	}
}

type GameJoinedIntgrEvent struct {
	Name       IntgrEventName     `json:"name"`
	LiveGameId string             `json:"liveGameId"`
	PlayerId   string             `json:"playerId"`
	Camera     viewmodel.CameraVm `json:"camera"`
	MapSize    viewmodel.SizeVm   `json:"mapSize"`
	View       viewmodel.ViewVm   `json:"view"`
}

func NewGameJoinedIntgrEvent(liveGameId string, playerId string, cameraVm viewmodel.CameraVm, mapSize viewmodel.SizeVm, viewVm viewmodel.ViewVm) GameJoinedIntgrEvent {
	return GameJoinedIntgrEvent{
		Name:       GameJoinedIntgrEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		Camera:     cameraVm,
		MapSize:    mapSize,
		View:       viewVm,
	}
}

type CameraChangedIntgrEvent struct {
	Name       IntgrEventName     `json:"name"`
	LiveGameId string             `json:"liveGameId"`
	PlayerId   string             `json:"playerId"`
	Camera     viewmodel.CameraVm `json:"camera"`
	View       viewmodel.ViewVm   `json:"view"`
}

func NewCameraChangedIntgrEvent(liveGameId string, playerId string, cameraVm viewmodel.CameraVm, viewVm viewmodel.ViewVm) CameraChangedIntgrEvent {
	return CameraChangedIntgrEvent{
		Name:       CameraChangedIntgrEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		Camera:     cameraVm,
		View:       viewVm,
	}
}

type ViewUpdatedIntgrEvent struct {
	Name       IntgrEventName     `json:"name"`
	LiveGameId string             `json:"liveGameId"`
	PlayerId   string             `json:"playerId"`
	Camera     viewmodel.CameraVm `json:"camera"`
	View       viewmodel.ViewVm   `json:"view"`
}

func NewViewUpdatedIntgrEvent(liveGameId string, playerId string, cameraVm viewmodel.CameraVm, viewVm viewmodel.ViewVm) ViewUpdatedIntgrEvent {
	return ViewUpdatedIntgrEvent{
		Name:       ViewUpdatedIntgrEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		Camera:     cameraVm,
		View:       viewVm,
	}
}

type ViewChangedIntgrEvent struct {
	Name       IntgrEventName   `json:"name"`
	LiveGameId string           `json:"liveGameId"`
	PlayerId   string           `json:"playerId"`
	View       viewmodel.ViewVm `json:"view"`
}

func NewViewChangedIntgrEvent(liveGameId string, playerId string, viewVm viewmodel.ViewVm) ViewChangedIntgrEvent {
	return ViewChangedIntgrEvent{
		Name:       ViewChangedIntgrEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		View:       viewVm,
	}
}

type LeaveGameRequestedIntgrEvent struct {
	Name       IntgrEventName `json:"name"`
	LiveGameId string         `json:"liveGameId"`
	PlayerId   string         `json:"playerId"`
}

func NewLeaveGameRequestedIntgrEvent(liveGameId string, playerId string) LeaveGameRequestedIntgrEvent {
	return LeaveGameRequestedIntgrEvent{
		Name:       LeaveGameRequestedIntgrEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
	}
}
