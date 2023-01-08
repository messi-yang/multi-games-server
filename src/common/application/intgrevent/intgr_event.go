package intgrevent

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
)

type IntgrEventName string

const (
	AddPlayerRequestedIntgrEventName  IntgrEventName = "ADD_PLAYER_REQUESTED"
	BuildItemRequestedIntgrEventName  IntgrEventName = "BUILD_ITEM_REQUESTED"
	DestroyItemRequestedEventName     IntgrEventName = "DESTROY_ITEM_REQUESTED"
	GameInfoUpdatedEventName          IntgrEventName = "GAME_INFO_UPDATED"
	MapRangeObservedEventName         IntgrEventName = "MAP_RANGE_OBSERVED"
	ObservedMapRangeUpdatedEventName  IntgrEventName = "OBSERVED_MAP_RANGE_UPDATED"
	ObserveMapRangeRequestedEventName IntgrEventName = "OBSERVE_MAP_RANGE_REQUESTED"
	RemovePlayerRequestedEventName    IntgrEventName = "REMOVE_PLAYER_REQUESTED"
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

type MapRangeObservedEvent struct {
	Name       IntgrEventName              `json:"name"`
	LiveGameId string                      `json:"liveGameId"`
	PlayerId   string                      `json:"playerId"`
	MapRange   viewmodel.MapRangeViewModel `json:"mapRange"`
	UnitMap    viewmodel.UnitMapViewModel  `json:"unitMap"`
}

func NewMapRangeObservedEvent(liveGameId string, playerId string, mapRange viewmodel.MapRangeViewModel, unitMap viewmodel.UnitMapViewModel) MapRangeObservedEvent {
	return MapRangeObservedEvent{
		Name:       MapRangeObservedEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		MapRange:   mapRange,
		UnitMap:    unitMap,
	}
}

type ObservedMapRangeUpdatedEvent struct {
	Name       IntgrEventName              `json:"name"`
	LiveGameId string                      `json:"liveGameId"`
	PlayerId   string                      `json:"playerId"`
	MapRange   viewmodel.MapRangeViewModel `json:"mapRange"`
	UnitMap    viewmodel.UnitMapViewModel  `json:"unitMap"`
}

func NewObservedMapRangeUpdatedEvent(liveGameId string, playerId string, mapRange viewmodel.MapRangeViewModel, unitMap viewmodel.UnitMapViewModel) ObservedMapRangeUpdatedEvent {
	return ObservedMapRangeUpdatedEvent{
		Name:       ObservedMapRangeUpdatedEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		MapRange:   mapRange,
		UnitMap:    unitMap,
	}
}

type ObserveMapRangeRequestedEvent struct {
	Name       IntgrEventName              `json:"name"`
	LiveGameId string                      `json:"liveGameId"`
	PlayerId   string                      `json:"playerId"`
	MapRange   viewmodel.MapRangeViewModel `json:"mapRange"`
}

func NewObserveMapRangeRequestedEvent(liveGameId string, playerId string, mapRange viewmodel.MapRangeViewModel) ObserveMapRangeRequestedEvent {
	return ObserveMapRangeRequestedEvent{
		Name:       ObserveMapRangeRequestedEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		MapRange:   mapRange,
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
