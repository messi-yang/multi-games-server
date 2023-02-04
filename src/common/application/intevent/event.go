package intevent

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
)

type IntEventName string

const (
	JoinGameRequestedIntEventName    IntEventName = "JOIN_GAME_REQUESTED"
	MoveRequestedIntEventName        IntEventName = "MOVE_REQUESTED"
	BuildItemRequestedIntEventName   IntEventName = "BUILD_ITEM_REQUESTED"
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

func NewJoinGameRequestedIntEvent(liveGameId string, playerId string) JoinGameRequestedIntEvent {
	return JoinGameRequestedIntEvent{
		Name:       JoinGameRequestedIntEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
	}
}

type MoveRequestedIntEvent struct {
	Name       IntEventName `json:"name"`
	LiveGameId string       `json:"liveGameId"`
	PlayerId   string       `json:"playerId"`
	Direction  int8         `json:"direction"`
}

func NewMoveRequestedIntEvent(liveGameId string, playerId string, direction int8) MoveRequestedIntEvent {
	return MoveRequestedIntEvent{
		Name:       MoveRequestedIntEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		Direction:  direction,
	}
}

type BuildItemRequestedIntEvent struct {
	Name       IntEventName         `json:"name"`
	LiveGameId string               `json:"liveGameId"`
	PlayerId   string               `json:"playerId"`
	Location   viewmodel.LocationVm `json:"location"`
	ItemId     string               `json:"itemId"`
}

func NewBuildItemRequestedIntEvent(liveGameId string, playerIdVm string, locationVm viewmodel.LocationVm, itemId string) BuildItemRequestedIntEvent {
	return BuildItemRequestedIntEvent{
		Name:       BuildItemRequestedIntEventName,
		LiveGameId: liveGameId,
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

func NewDestroyItemRequestedIntEvent(liveGameId string, playerIdVm string, locationVm viewmodel.LocationVm) DestroyItemRequestedIntEvent {
	return DestroyItemRequestedIntEvent{
		Name:       DestroyItemRequestedIntEventName,
		LiveGameId: liveGameId,
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

func NewGameJoinedIntEvent(liveGameId string, playerVms []viewmodel.PlayerVm, mapSize viewmodel.SizeVm, viewVm viewmodel.ViewVm) GameJoinedIntEvent {
	return GameJoinedIntEvent{
		Name:       GameJoinedIntEventName,
		LiveGameId: liveGameId,
		Players:    playerVms,
		MapSize:    mapSize,
		View:       viewVm,
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
