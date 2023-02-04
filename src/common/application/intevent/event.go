package intevent

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
)

type IntEventName string

const (
	JoinGameRequestedIntEventName    IntEventName = "JOIN_GAME_REQUESTED"
	MoveRequestedIntEventName        IntEventName = "MOVE_REQUESTED"
	PlaceItemRequestedIntEventName   IntEventName = "PLACE_ITEM_REQUESTED"
	DestroyItemRequestedIntEventName IntEventName = "DESTROY_ITEM_REQUESTED"
	LeaveGameRequestedIntEventName   IntEventName = "LEAVE_GAME_REQUESTED"
	GameJoinedIntEventName           IntEventName = "GAME_JOINED"
	PlayersUpdatedIntEventName       IntEventName = "PLAYERS_UPDATED"
	ViewUpdatedIntEventName          IntEventName = "VIEW_UPDATED"
)

type GenericIntEvent struct {
	Name IntEventName `json:"name"`
}

type JoinGameRequestedIntEvent struct {
	Name       IntEventName `json:"name"`
	LiveGameId string       `json:"liveGameId"`
	PlayerId   string       `json:"playerId"`
}

func NewJoinGameRequestedIntEvent(liveGameIdVm string, playerIdVm string) JoinGameRequestedIntEvent {
	return JoinGameRequestedIntEvent{
		Name:       JoinGameRequestedIntEventName,
		LiveGameId: liveGameIdVm,
		PlayerId:   playerIdVm,
	}
}

type MoveRequestedIntEvent struct {
	Name       IntEventName `json:"name"`
	LiveGameId string       `json:"liveGameId"`
	PlayerId   string       `json:"playerId"`
	Direction  int8         `json:"direction"`
}

func NewMoveRequestedIntEvent(liveGameIdVm string, playerIdVm string, direction int8) MoveRequestedIntEvent {
	return MoveRequestedIntEvent{
		Name:       MoveRequestedIntEventName,
		LiveGameId: liveGameIdVm,
		PlayerId:   playerIdVm,
		Direction:  direction,
	}
}

type PlaceItemRequestedIntEvent struct {
	Name       IntEventName         `json:"name"`
	LiveGameId string               `json:"liveGameId"`
	PlayerId   string               `json:"playerId"`
	Location   viewmodel.LocationVm `json:"location"`
	ItemId     string               `json:"itemId"`
}

func NewPlaceItemRequestedIntEvent(liveGameIdVm string, playerIdVm string, locationVm viewmodel.LocationVm, itemId string) PlaceItemRequestedIntEvent {
	return PlaceItemRequestedIntEvent{
		Name:       PlaceItemRequestedIntEventName,
		LiveGameId: liveGameIdVm,
		PlayerId:   playerIdVm,
		Location:   locationVm,
		ItemId:     itemId,
	}
}

type DestroyItemRequestedIntEvent struct {
	Name       IntEventName         `json:"name"`
	LiveGameId string               `json:"liveGameId"`
	PlayerId   string               `json:"playerId"`
	Location   viewmodel.LocationVm `json:"location"`
}

func NewDestroyItemRequestedIntEvent(liveGameIdVm string, playerIdVm string, locationVm viewmodel.LocationVm) DestroyItemRequestedIntEvent {
	return DestroyItemRequestedIntEvent{
		Name:       DestroyItemRequestedIntEventName,
		LiveGameId: liveGameIdVm,
		PlayerId:   playerIdVm,
		Location:   locationVm,
	}
}

type GameJoinedIntEvent struct {
	Name       IntEventName         `json:"name"`
	LiveGameId string               `json:"liveGameId"`
	Players    []viewmodel.PlayerVm `json:"players"`
	MapSize    viewmodel.SizeVm     `json:"mapSize"`
	View       viewmodel.ViewVm     `json:"view"`
}

func NewGameJoinedIntEvent(liveGameIdVm string, playerVms []viewmodel.PlayerVm, mapSizeVm viewmodel.SizeVm, viewVm viewmodel.ViewVm) GameJoinedIntEvent {
	return GameJoinedIntEvent{
		Name:       GameJoinedIntEventName,
		LiveGameId: liveGameIdVm,
		Players:    playerVms,
		MapSize:    mapSizeVm,
		View:       viewVm,
	}
}

type PlayersUpdatedIntEvent struct {
	Name       IntEventName         `json:"name"`
	LiveGameId string               `json:"liveGameId"`
	Players    []viewmodel.PlayerVm `json:"players"`
}

func NewPlayersUpdatedIntEvent(liveGameIdVm string, playerVms []viewmodel.PlayerVm) PlayersUpdatedIntEvent {
	return PlayersUpdatedIntEvent{
		Name:       PlayersUpdatedIntEventName,
		LiveGameId: liveGameIdVm,
		Players:    playerVms,
	}
}

type ViewUpdatedIntEvent struct {
	Name       IntEventName     `json:"name"`
	LiveGameId string           `json:"liveGameId"`
	PlayerId   string           `json:"playerId"`
	View       viewmodel.ViewVm `json:"view"`
}

func NewViewUpdatedIntEvent(liveGameIdVm string, playerIdVm string, viewVm viewmodel.ViewVm) ViewUpdatedIntEvent {
	return ViewUpdatedIntEvent{
		Name:       ViewUpdatedIntEventName,
		LiveGameId: liveGameIdVm,
		PlayerId:   playerIdVm,
		View:       viewVm,
	}
}

type LeaveGameRequestedIntEvent struct {
	Name       IntEventName `json:"name"`
	LiveGameId string       `json:"liveGameId"`
	PlayerId   string       `json:"playerId"`
}

func NewLeaveGameRequestedIntEvent(liveGameIdVm string, playerIdVm string) LeaveGameRequestedIntEvent {
	return LeaveGameRequestedIntEvent{
		Name:       LeaveGameRequestedIntEventName,
		LiveGameId: liveGameIdVm,
		PlayerId:   playerIdVm,
	}
}
