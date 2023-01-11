package intgrevent

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
)

type IntgrEventName string

const (
	JoinLiveGameRequestedIntgrEventName  IntgrEventName = "ADD_PLAYER_REQUESTED"
	BuildItemRequestedIntgrEventName     IntgrEventName = "BUILD_ITEM_REQUESTED"
	DestroyItemRequestedIntgrEventName   IntgrEventName = "DESTROY_ITEM_REQUESTED"
	GameJoinedIntgrEventName             IntgrEventName = "GAME_JOINED"
	RangeObservedIntgrEventName          IntgrEventName = "RANGE_OBSERVED"
	ObservedRangeUpdatedIntgrEventName   IntgrEventName = "OBSERVED_RANGE_UPDATED"
	ObserveRangeRequestedIntgrEventName  IntgrEventName = "OBSERVE_RANGE_REQUESTED"
	LeaveLiveGameRequestedIntgrEventName IntgrEventName = "REMOVE_PLAYER_REQUESTED"
)

type GenericIntgrEvent struct {
	Name IntgrEventName `json:"name"`
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

type JoinLiveGameRequestedIntgrEvent struct {
	Name       IntgrEventName `json:"name"`
	LiveGameId string         `json:"liveGameId"`
	PlayerId   string         `json:"playerId"`
}

func NewJoinLiveGameRequestedIntgrEvent(liveGameId string, playerId string) JoinLiveGameRequestedIntgrEvent {
	return JoinLiveGameRequestedIntgrEvent{
		Name:       JoinLiveGameRequestedIntgrEventName,
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
	Range      viewmodel.RangeVm     `json:"range"`
	Map        viewmodel.MapVm       `json:"map"`
}

func NewGameJoinedIntgrEvent(liveGameId string, playerId string, cameraVm viewmodel.CameraVm, dimensionVm viewmodel.DimensionVm, rangeVm viewmodel.RangeVm, mapVm viewmodel.MapVm) GameJoinedIntgrEvent {
	return GameJoinedIntgrEvent{
		Name:       GameJoinedIntgrEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		Camera:     cameraVm,
		Dimension:  dimensionVm,
		Range:      rangeVm,
		Map:        mapVm,
	}
}

type RangeObservedIntgrEvent struct {
	Name       IntgrEventName    `json:"name"`
	LiveGameId string            `json:"liveGameId"`
	PlayerId   string            `json:"playerId"`
	Range      viewmodel.RangeVm `json:"range"`
	Map        viewmodel.MapVm   `json:"map"`
}

func NewRangeObservedIntgrEvent(liveGameId string, playerId string, rangeVm viewmodel.RangeVm, mapVm viewmodel.MapVm) RangeObservedIntgrEvent {
	return RangeObservedIntgrEvent{
		Name:       RangeObservedIntgrEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		Range:      rangeVm,
		Map:        mapVm,
	}
}

type ObservedRangeUpdatedIntgrEvent struct {
	Name       IntgrEventName    `json:"name"`
	LiveGameId string            `json:"liveGameId"`
	PlayerId   string            `json:"playerId"`
	Range      viewmodel.RangeVm `json:"range"`
	Map        viewmodel.MapVm   `json:"map"`
}

func NewObservedRangeUpdatedIntgrEvent(liveGameId string, playerId string, rangeVm viewmodel.RangeVm, mapVm viewmodel.MapVm) ObservedRangeUpdatedIntgrEvent {
	return ObservedRangeUpdatedIntgrEvent{
		Name:       ObservedRangeUpdatedIntgrEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		Range:      rangeVm,
		Map:        mapVm,
	}
}

type ObserveRangeRequestedIntgrEvent struct {
	Name       IntgrEventName    `json:"name"`
	LiveGameId string            `json:"liveGameId"`
	PlayerId   string            `json:"playerId"`
	Range      viewmodel.RangeVm `json:"range"`
}

func NewObserveRangeRequestedIntgrEvent(liveGameId string, playerId string, rangeVm viewmodel.RangeVm) ObserveRangeRequestedIntgrEvent {
	return ObserveRangeRequestedIntgrEvent{
		Name:       ObserveRangeRequestedIntgrEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		Range:      rangeVm,
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
