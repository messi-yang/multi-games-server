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
	Name     IntEventName `json:"name"`
	GameId   string       `json:"gameId"`
	PlayerId string       `json:"playerId"`
}

func NewJoinGameRequestedIntEvent(gameIdVm string, playerIdVm string) JoinGameRequestedIntEvent {
	return JoinGameRequestedIntEvent{
		Name:     JoinGameRequestedIntEventName,
		GameId:   gameIdVm,
		PlayerId: playerIdVm,
	}
}

type MoveRequestedIntEvent struct {
	Name      IntEventName `json:"name"`
	GameId    string       `json:"gameId"`
	PlayerId  string       `json:"playerId"`
	Direction int8         `json:"direction"`
}

func NewMoveRequestedIntEvent(gameIdVm string, playerIdVm string, direction int8) MoveRequestedIntEvent {
	return MoveRequestedIntEvent{
		Name:      MoveRequestedIntEventName,
		GameId:    gameIdVm,
		PlayerId:  playerIdVm,
		Direction: direction,
	}
}

type PlaceItemRequestedIntEvent struct {
	Name     IntEventName         `json:"name"`
	GameId   string               `json:"gameId"`
	PlayerId string               `json:"playerId"`
	Location viewmodel.LocationVm `json:"location"`
	ItemId   string               `json:"itemId"`
}

func NewPlaceItemRequestedIntEvent(gameIdVm string, playerIdVm string, locationVm viewmodel.LocationVm, itemId string) PlaceItemRequestedIntEvent {
	return PlaceItemRequestedIntEvent{
		Name:     PlaceItemRequestedIntEventName,
		GameId:   gameIdVm,
		PlayerId: playerIdVm,
		Location: locationVm,
		ItemId:   itemId,
	}
}

type DestroyItemRequestedIntEvent struct {
	Name     IntEventName         `json:"name"`
	GameId   string               `json:"gameId"`
	PlayerId string               `json:"playerId"`
	Location viewmodel.LocationVm `json:"location"`
}

func NewDestroyItemRequestedIntEvent(gameIdVm string, playerIdVm string, locationVm viewmodel.LocationVm) DestroyItemRequestedIntEvent {
	return DestroyItemRequestedIntEvent{
		Name:     DestroyItemRequestedIntEventName,
		GameId:   gameIdVm,
		PlayerId: playerIdVm,
		Location: locationVm,
	}
}

type GameJoinedIntEvent struct {
	Name    IntEventName         `json:"name"`
	GameId  string               `json:"gameId"`
	Players []viewmodel.PlayerVm `json:"players"`
	MapSize viewmodel.SizeVm     `json:"mapSize"`
	View    viewmodel.ViewVm     `json:"view"`
}

func NewGameJoinedIntEvent(gameIdVm string, playerVms []viewmodel.PlayerVm, mapSizeVm viewmodel.SizeVm, viewVm viewmodel.ViewVm) GameJoinedIntEvent {
	return GameJoinedIntEvent{
		Name:    GameJoinedIntEventName,
		GameId:  gameIdVm,
		Players: playerVms,
		MapSize: mapSizeVm,
		View:    viewVm,
	}
}

type PlayersUpdatedIntEvent struct {
	Name    IntEventName         `json:"name"`
	GameId  string               `json:"gameId"`
	Players []viewmodel.PlayerVm `json:"players"`
}

func NewPlayersUpdatedIntEvent(gameIdVm string, playerVms []viewmodel.PlayerVm) PlayersUpdatedIntEvent {
	return PlayersUpdatedIntEvent{
		Name:    PlayersUpdatedIntEventName,
		GameId:  gameIdVm,
		Players: playerVms,
	}
}

type ViewUpdatedIntEvent struct {
	Name     IntEventName     `json:"name"`
	GameId   string           `json:"gameId"`
	PlayerId string           `json:"playerId"`
	View     viewmodel.ViewVm `json:"view"`
}

func NewViewUpdatedIntEvent(gameIdVm string, playerIdVm string, viewVm viewmodel.ViewVm) ViewUpdatedIntEvent {
	return ViewUpdatedIntEvent{
		Name:     ViewUpdatedIntEventName,
		GameId:   gameIdVm,
		PlayerId: playerIdVm,
		View:     viewVm,
	}
}

type LeaveGameRequestedIntEvent struct {
	Name     IntEventName `json:"name"`
	GameId   string       `json:"gameId"`
	PlayerId string       `json:"playerId"`
}

func NewLeaveGameRequestedIntEvent(gameIdVm string, playerIdVm string) LeaveGameRequestedIntEvent {
	return LeaveGameRequestedIntEvent{
		Name:     LeaveGameRequestedIntEventName,
		GameId:   gameIdVm,
		PlayerId: playerIdVm,
	}
}
