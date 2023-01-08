package intgrevent

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
)

type IntgrEventName string

const (
	AddPlayerRequestedIntgrEventName    IntgrEventName = "ADD_PLAYER_REQUESTED"
	BuildItemRequestedIntgrEventName    IntgrEventName = "BUILD_ITEM_REQUESTED"
	DestroyItemRequestedIntgrEventName  IntgrEventName = "DESTROY_ITEM_REQUESTED"
	GameInfoUpdatedIntgrEventName       IntgrEventName = "GAME_INFO_UPDATED"
	RangeObservedIntgrEventName         IntgrEventName = "RANGE_OBSERVED"
	ObservedRangeUpdatedIntgrEventName  IntgrEventName = "OBSERVED_RANGE_UPDATED"
	ObserveRangeRequestedIntgrEventName IntgrEventName = "OBSERVE_RANGE_REQUESTED"
	RemovePlayerRequestedIntgrEventName IntgrEventName = "REMOVE_PLAYER_REQUESTED"
)

type GenericIntgrEvent struct {
	Name IntgrEventName `json:"name"`
}

type BuildItemRequestedIntgrEvent struct {
	Name       IntgrEventName     `json:"name"`
	LiveGameId string             `json:"liveGameId"`
	Location   viewmodel.Location `json:"location"`
	ItemId     string             `json:"locations"`
}

func NewBuildItemRequestedIntgrEvent(liveGameId string, location viewmodel.Location, itemId string) BuildItemRequestedIntgrEvent {
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

type DestroyItemRequestedIntgrEvent struct {
	Name       IntgrEventName     `json:"name"`
	LiveGameId string             `json:"liveGameId"`
	Location   viewmodel.Location `json:"location"`
}

func NewDestroyItemRequestedIntgrEvent(liveGameId string, location viewmodel.Location) DestroyItemRequestedIntgrEvent {
	return DestroyItemRequestedIntgrEvent{
		Name:       DestroyItemRequestedIntgrEventName,
		LiveGameId: liveGameId,
		Location:   location,
	}
}

type GameInfoUpdatedIntgrEvent struct {
	Name       IntgrEventName    `json:"name"`
	LiveGameId string            `json:"liveGameId"`
	PlayerId   string            `json:"playerId"`
	MapSize    viewmodel.MapSize `json:"mapSize"`
}

func NewGameInfoUpdatedIntgrEvent(liveGameId string, playerId string, mapSize viewmodel.MapSize) GameInfoUpdatedIntgrEvent {
	return GameInfoUpdatedIntgrEvent{
		Name:       GameInfoUpdatedIntgrEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		MapSize:    mapSize,
	}
}

type RangeObservedIntgrEvent struct {
	Name       IntgrEventName    `json:"name"`
	LiveGameId string            `json:"liveGameId"`
	PlayerId   string            `json:"playerId"`
	Range      viewmodel.Range   `json:"range"`
	UnitMap    viewmodel.UnitMap `json:"unitMap"`
}

func NewRangeObservedIntgrEvent(liveGameId string, playerId string, rangeVm viewmodel.Range, unitMap viewmodel.UnitMap) RangeObservedIntgrEvent {
	return RangeObservedIntgrEvent{
		Name:       RangeObservedIntgrEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		Range:      rangeVm,
		UnitMap:    unitMap,
	}
}

type ObservedRangeUpdatedIntgrEvent struct {
	Name       IntgrEventName    `json:"name"`
	LiveGameId string            `json:"liveGameId"`
	PlayerId   string            `json:"playerId"`
	Range      viewmodel.Range   `json:"range"`
	UnitMap    viewmodel.UnitMap `json:"unitMap"`
}

func NewObservedRangeUpdatedIntgrEvent(liveGameId string, playerId string, rangeVm viewmodel.Range, unitMap viewmodel.UnitMap) ObservedRangeUpdatedIntgrEvent {
	return ObservedRangeUpdatedIntgrEvent{
		Name:       ObservedRangeUpdatedIntgrEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		Range:      rangeVm,
		UnitMap:    unitMap,
	}
}

type ObserveRangeRequestedIntgrEvent struct {
	Name       IntgrEventName  `json:"name"`
	LiveGameId string          `json:"liveGameId"`
	PlayerId   string          `json:"playerId"`
	Range      viewmodel.Range `json:"range"`
}

func NewObserveRangeRequestedIntgrEvent(liveGameId string, playerId string, rangeVm viewmodel.Range) ObserveRangeRequestedIntgrEvent {
	return ObserveRangeRequestedIntgrEvent{
		Name:       ObserveRangeRequestedIntgrEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		Range:      rangeVm,
	}
}

type RemovePlayerRequestedIntgrEvent struct {
	Name       IntgrEventName `json:"name"`
	LiveGameId string         `json:"liveGameId"`
	PlayerId   string         `json:"playerId"`
}

func NewRemovePlayerRequestedIntgrEvent(liveGameId string, playerId string) RemovePlayerRequestedIntgrEvent {
	return RemovePlayerRequestedIntgrEvent{
		Name:       RemovePlayerRequestedIntgrEventName,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
	}
}
