package intevent

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
)

type intEventName string

const (
	ChangeCameraRequestedintEventName intEventName = "CHANGE_CAMERA_REQUESTED"
	JoinGameRequestedintEventName     intEventName = "JOIN_GAME_REQUESTED"
	BuildItemRequestedintEventName    intEventName = "BUILD_ITEM_REQUESTED"
	DestroyItemRequestedintEventName  intEventName = "DESTROY_ITEM_REQUESTED"
	LeaveGameRequestedintEventName    intEventName = "LEAVE_GAME_REQUESTED"
	GameJoinedintEventName            intEventName = "GAME_JOINED"
	CameraChangedintEventName         intEventName = "CAMERA_CHANGED"
	ViewUpdatedintEventName           intEventName = "VIEW_UPDATED"
	ViewChangedintEventName           intEventName = "VIEW_CHANGED"
)

type GenericintEvent struct {
	Name intEventName `json:"name"`
}
type ChangeCameraRequestedintEvent struct {
	Name       intEventName       `json:"name"`
	LiveGameId string             `json:"liveGameId"`
	PlayerId   string             `json:"playerId"`
	Camera     viewmodel.CameraVm `json:"camera"`
}

func NewChangeCameraRequestedintEvent(liveGameId string, playerId string, cameraVm viewmodel.CameraVm) ChangeCameraRequestedintEvent {
	return ChangeCameraRequestedintEvent{
		Name:       ChangeCameraRequestedintEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		Camera:     cameraVm,
	}
}

type BuildItemRequestedintEvent struct {
	Name       intEventName         `json:"name"`
	LiveGameId string               `json:"liveGameId"`
	Location   viewmodel.LocationVm `json:"location"`
	ItemId     string               `json:"itemId"`
}

func NewBuildItemRequestedintEvent(liveGameId string, locationVm viewmodel.LocationVm, itemId string) BuildItemRequestedintEvent {
	return BuildItemRequestedintEvent{
		Name:       BuildItemRequestedintEventName,
		LiveGameId: liveGameId,
		Location:   locationVm,
		ItemId:     itemId,
	}
}

type JoinGameRequestedintEvent struct {
	Name       intEventName `json:"name"`
	LiveGameId string       `json:"liveGameId"`
	PlayerId   string       `json:"playerId"`
}

func NewJoinGameRequestedintEvent(liveGameId string, playerId string) JoinGameRequestedintEvent {
	return JoinGameRequestedintEvent{
		Name:       JoinGameRequestedintEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
	}
}

type DestroyItemRequestedintEvent struct {
	Name       intEventName         `json:"name"`
	LiveGameId string               `json:"liveGameId"`
	Location   viewmodel.LocationVm `json:"location"`
}

func NewDestroyItemRequestedintEvent(liveGameId string, locationVm viewmodel.LocationVm) DestroyItemRequestedintEvent {
	return DestroyItemRequestedintEvent{
		Name:       DestroyItemRequestedintEventName,
		LiveGameId: liveGameId,
		Location:   locationVm,
	}
}

type GameJoinedintEvent struct {
	Name       intEventName       `json:"name"`
	LiveGameId string             `json:"liveGameId"`
	PlayerId   string             `json:"playerId"`
	Camera     viewmodel.CameraVm `json:"camera"`
	MapSize    viewmodel.SizeVm   `json:"mapSize"`
	View       viewmodel.ViewVm   `json:"view"`
}

func NewGameJoinedintEvent(liveGameId string, playerId string, cameraVm viewmodel.CameraVm, mapSize viewmodel.SizeVm, viewVm viewmodel.ViewVm) GameJoinedintEvent {
	return GameJoinedintEvent{
		Name:       GameJoinedintEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		Camera:     cameraVm,
		MapSize:    mapSize,
		View:       viewVm,
	}
}

type CameraChangedintEvent struct {
	Name       intEventName       `json:"name"`
	LiveGameId string             `json:"liveGameId"`
	PlayerId   string             `json:"playerId"`
	Camera     viewmodel.CameraVm `json:"camera"`
}

func NewCameraChangedintEvent(liveGameId string, playerId string, cameraVm viewmodel.CameraVm) CameraChangedintEvent {
	return CameraChangedintEvent{
		Name:       CameraChangedintEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		Camera:     cameraVm,
	}
}

type ViewUpdatedintEvent struct {
	Name       intEventName     `json:"name"`
	LiveGameId string           `json:"liveGameId"`
	PlayerId   string           `json:"playerId"`
	View       viewmodel.ViewVm `json:"view"`
}

func NewViewUpdatedintEvent(liveGameId string, playerId string, viewVm viewmodel.ViewVm) ViewUpdatedintEvent {
	return ViewUpdatedintEvent{
		Name:       ViewUpdatedintEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		View:       viewVm,
	}
}

type ViewChangedintEvent struct {
	Name       intEventName     `json:"name"`
	LiveGameId string           `json:"liveGameId"`
	PlayerId   string           `json:"playerId"`
	View       viewmodel.ViewVm `json:"view"`
}

func NewViewChangedintEvent(liveGameId string, playerId string, viewVm viewmodel.ViewVm) ViewChangedintEvent {
	return ViewChangedintEvent{
		Name:       ViewChangedintEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		View:       viewVm,
	}
}

type LeaveGameRequestedintEvent struct {
	Name       intEventName `json:"name"`
	LiveGameId string       `json:"liveGameId"`
	PlayerId   string       `json:"playerId"`
}

func NewLeaveGameRequestedintEvent(liveGameId string, playerId string) LeaveGameRequestedintEvent {
	return LeaveGameRequestedintEvent{
		Name:       LeaveGameRequestedintEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
	}
}
