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
	RangeObservedEventName           IntgrEventName = "RANGE_OBSERVED"
	ObservedRangeUpdatedEventName    IntgrEventName = "OBSERVED_RANGE_UPDATED"
	ObserveRangeRequestedEventName   IntgrEventName = "OBSERVE_RANGE_REQUESTED"
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

type RangeObservedEvent struct {
	Name       IntgrEventName             `json:"name"`
	LiveGameId string                     `json:"liveGameId"`
	PlayerId   string                     `json:"playerId"`
	Range      viewmodel.RangeViewModel   `json:"range"`
	UnitMap    viewmodel.UnitMapViewModel `json:"unitMap"`
}

func NewRangeObservedEvent(liveGameId string, playerId string, rangeVm viewmodel.RangeViewModel, unitMap viewmodel.UnitMapViewModel) RangeObservedEvent {
	return RangeObservedEvent{
		Name:       RangeObservedEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		Range:      rangeVm,
		UnitMap:    unitMap,
	}
}

type ObservedRangeUpdatedEvent struct {
	Name       IntgrEventName             `json:"name"`
	LiveGameId string                     `json:"liveGameId"`
	PlayerId   string                     `json:"playerId"`
	Range      viewmodel.RangeViewModel   `json:"range"`
	UnitMap    viewmodel.UnitMapViewModel `json:"unitMap"`
}

func NewObservedRangeUpdatedEvent(liveGameId string, playerId string, rangeVm viewmodel.RangeViewModel, unitMap viewmodel.UnitMapViewModel) ObservedRangeUpdatedEvent {
	return ObservedRangeUpdatedEvent{
		Name:       ObservedRangeUpdatedEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		Range:      rangeVm,
		UnitMap:    unitMap,
	}
}

type ObserveRangeRequestedEvent struct {
	Name       IntgrEventName           `json:"name"`
	LiveGameId string                   `json:"liveGameId"`
	PlayerId   string                   `json:"playerId"`
	Range      viewmodel.RangeViewModel `json:"range"`
}

func NewObserveRangeRequestedEvent(liveGameId string, playerId string, rangeVm viewmodel.RangeViewModel) ObserveRangeRequestedEvent {
	return ObserveRangeRequestedEvent{
		Name:       ObserveRangeRequestedEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		Range:      rangeVm,
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
