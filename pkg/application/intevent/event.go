package intevent

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/viewmodel"
)

type IntEventName string

const (
	PlaceItemRequestedIntEventName   IntEventName = "PLACE_ITEM_REQUESTED"
	DestroyItemRequestedIntEventName IntEventName = "DESTROY_ITEM_REQUESTED"
	PlayersUpdatedIntEventName       IntEventName = "PLAYERS_UPDATED"
	ViewUpdatedIntEventName          IntEventName = "VIEW_UPDATED"
)

type GenericIntEvent struct {
	Name IntEventName `json:"name"`
}

type PlaceItemRequestedIntEvent struct {
	Name     IntEventName         `json:"name"`
	GameId   string               `json:"gameId"`
	PlayerId string               `json:"playerId"`
	Location viewmodel.LocationVm `json:"location"`
	ItemId   int16                `json:"itemId"`
}

func NewPlaceItemRequestedIntEvent(gameIdVm string, playerIdVm string, locationVm viewmodel.LocationVm, itemIdVm int16) PlaceItemRequestedIntEvent {
	return PlaceItemRequestedIntEvent{
		Name:     PlaceItemRequestedIntEventName,
		GameId:   gameIdVm,
		PlayerId: playerIdVm,
		Location: locationVm,
		ItemId:   itemIdVm,
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
