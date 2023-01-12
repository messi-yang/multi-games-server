package intgrevent

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
)

type IntgrEventName string

const (
	ChangeCameraRequestedIntgrEventName  IntgrEventName = "CHANGE_CAMERA_REQUESTED"
	JoinGameRequestedIntgrEventName      IntgrEventName = "JOIN_GAME_REQUESTED"
	BuildItemRequestedIntgrEventName     IntgrEventName = "BUILD_ITEM_REQUESTED"
	DestroyItemRequestedIntgrEventName   IntgrEventName = "DESTROY_ITEM_REQUESTED"
	LeaveLiveGameRequestedIntgrEventName IntgrEventName = "REMOVE_PLAYER_REQUESTED"
	GameJoinedIntgrEventName             IntgrEventName = "GAME_JOINED"
	CameraChangedIntgrEventName          IntgrEventName = "CAMERA_CHANGED"
	ViewUpdatedIntgrEventName            IntgrEventName = "VIEW_UPDATED"
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
	Name       IntgrEventName        `json:"name"`
	LiveGameId string                `json:"liveGameId"`
	PlayerId   string                `json:"playerId"`
	Camera     viewmodel.CameraVm    `json:"camera"`
	Dimension  viewmodel.DimensionVm `json:"dimension"`
	View       viewmodel.ViewVm      `json:"view"`
}

func NewGameJoinedIntgrEvent(liveGameId string, playerId string, cameraVm viewmodel.CameraVm, dimensionVm viewmodel.DimensionVm, viewVm viewmodel.ViewVm) GameJoinedIntgrEvent {
	return GameJoinedIntgrEvent{
		Name:       GameJoinedIntgrEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		Camera:     cameraVm,
		Dimension:  dimensionVm,
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

func NewViewUpdatedIntgrEvent(liveGameId string, playerId string, cameraVm viewmodel.CameraVm, viewVm viewmodel.ViewVm) GameJoinedIntgrEvent {
	return GameJoinedIntgrEvent{
		Name:       ViewUpdatedIntgrEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		Camera:     cameraVm,
		View:       viewVm,
	}
}

type LeaveLiveGameRequestedIntgrEvent struct {
	Name       IntgrEventName `json:"name"`
	LiveGameId string         `json:"liveGameId"`
	PlayerId   string         `json:"playerId"`
}

func NewLeaveLiveGameRequestedIntgrEvent(liveGameId string, playerId string) LeaveLiveGameRequestedIntgrEvent {
	return LeaveLiveGameRequestedIntgrEvent{
		Name:       LeaveLiveGameRequestedIntgrEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
	}
}
