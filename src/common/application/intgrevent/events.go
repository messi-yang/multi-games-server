package intgrevent

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
)

type IntgrEventName string

const (
	AddPlayerRequestedIntgrEventName IntgrEventName = "ADD_PLAYER_REQUESTED"
	BuildItemRequestedIntgrEventName IntgrEventName = "BUILD_ITEM_REQUESTED"
	DestroyItemRequestedEventName    IntgrEventName = "DESTROY_ITEM_REQUESTED"
	GameInfoUpdatedEventName         IntgrEventName = "GAME_INFO_UPDATED"
	ExtentObservedEventName          IntgrEventName = "MAP_RANGE_OBSERVED"
	ObservedExtentUpdatedEventName   IntgrEventName = "OBSERVED_MAP_RANGE_UPDATED"
	ObserveExtentRequestedEventName  IntgrEventName = "OBSERVE_MAP_RANGE_REQUESTED"
	RemovePlayerRequestedEventName   IntgrEventName = "REMOVE_PLAYER_REQUESTED"
)

type GenericIntgrEvent struct {
	Name IntgrEventName `json:"name"`
}

type BuildItemRequestedIntgrEvent struct {
	Name       IntgrEventName              `json:"name"`
	LiveGameId string                      `json:"liveGameId"`
	Location   viewmodel.LocationViewModel `json:"location"`
	ItemId     string                      `json:"locations"`
}

func NewBuildItemRequestedIntgrEvent(liveGameId string, location viewmodel.LocationViewModel, itemId string) BuildItemRequestedIntgrEvent {
	return BuildItemRequestedIntgrEvent{
		Name:       BuildItemRequestedIntgrEventName,
		LiveGameId: liveGameId,
		Location:   location,
		ItemId:     itemId,
	}
}

type AddPlayerRequestedIntgrEvent struct {
	Name       IntgrEventName `json:"name"`
	LiveGameId string         `json:"liveGameId"`
	PlayerId   string         `json:"playerId"`
}

func NewAddPlayerRequestedIntgrEvent(liveGameId string, playerId string) AddPlayerRequestedIntgrEvent {
	return AddPlayerRequestedIntgrEvent{
		Name:       AddPlayerRequestedIntgrEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
	}
}

type DestroyItemRequestedEvent struct {
	Name       IntgrEventName              `json:"name"`
	LiveGameId string                      `json:"liveGameId"`
	Location   viewmodel.LocationViewModel `json:"location"`
}

func NewDestroyItemRequested(liveGameId string, location viewmodel.LocationViewModel) DestroyItemRequestedEvent {
	return DestroyItemRequestedEvent{
		Name:       DestroyItemRequestedEventName,
		LiveGameId: liveGameId,
		Location:   location,
	}
}

type GameInfoUpdatedEvent struct {
	Name       IntgrEventName             `json:"name"`
	LiveGameId string                     `json:"liveGameId"`
	PlayerId   string                     `json:"playerId"`
	MapSize    viewmodel.MapSizeViewModel `json:"mapSize"`
}

func NewGameInfoUpdatedEvent(liveGameId string, playerId string, mapSize viewmodel.MapSizeViewModel) GameInfoUpdatedEvent {
	return GameInfoUpdatedEvent{
		Name:       GameInfoUpdatedEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		MapSize:    mapSize,
	}
}

type ExtentObservedEvent struct {
	Name       IntgrEventName             `json:"name"`
	LiveGameId string                     `json:"liveGameId"`
	PlayerId   string                     `json:"playerId"`
	Extent     viewmodel.ExtentViewModel  `json:"extent"`
	UnitMap    viewmodel.UnitMapViewModel `json:"unitMap"`
}

func NewExtentObservedEvent(liveGameId string, playerId string, extent viewmodel.ExtentViewModel, unitMap viewmodel.UnitMapViewModel) ExtentObservedEvent {
	return ExtentObservedEvent{
		Name:       ExtentObservedEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		Extent:     extent,
		UnitMap:    unitMap,
	}
}

type ObservedExtentUpdatedEvent struct {
	Name       IntgrEventName             `json:"name"`
	LiveGameId string                     `json:"liveGameId"`
	PlayerId   string                     `json:"playerId"`
	Extent     viewmodel.ExtentViewModel  `json:"extent"`
	UnitMap    viewmodel.UnitMapViewModel `json:"unitMap"`
}

func NewObservedExtentUpdatedEvent(liveGameId string, playerId string, extent viewmodel.ExtentViewModel, unitMap viewmodel.UnitMapViewModel) ObservedExtentUpdatedEvent {
	return ObservedExtentUpdatedEvent{
		Name:       ObservedExtentUpdatedEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		Extent:     extent,
		UnitMap:    unitMap,
	}
}

type ObserveExtentRequestedEvent struct {
	Name       IntgrEventName            `json:"name"`
	LiveGameId string                    `json:"liveGameId"`
	PlayerId   string                    `json:"playerId"`
	Extent     viewmodel.ExtentViewModel `json:"extent"`
}

func NewObserveExtentRequestedEvent(liveGameId string, playerId string, extent viewmodel.ExtentViewModel) ObserveExtentRequestedEvent {
	return ObserveExtentRequestedEvent{
		Name:       ObserveExtentRequestedEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		Extent:     extent,
	}
}

type RemovePlayerRequestedEvent struct {
	Name       IntgrEventName `json:"name"`
	LiveGameId string         `json:"liveGameId"`
	PlayerId   string         `json:"playerId"`
}

func NewRemovePlayerRequestedEvent(liveGameId string, playerId string) RemovePlayerRequestedEvent {
	return RemovePlayerRequestedEvent{
		Name:       RemovePlayerRequestedEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
	}
}
