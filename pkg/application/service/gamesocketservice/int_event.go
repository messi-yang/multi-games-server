package gamesocketservice

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/viewmodel"
)

func CreateGameIntEventChannel(gameId string) string {
	return fmt.Sprintf("GAME_%s", gameId)
}

type IntEventName string

const (
	PlayerUpdatedIntEventName IntEventName = "PLAYER_UPDATED"
	UnitUpdatedIntEventName   IntEventName = "UNIT_UPDATED"
)

type GenericIntEvent struct {
	Name IntEventName `json:"name"`
}

type PlayerUpdatedIntEvent struct {
	Name     IntEventName `json:"name"`
	GameId   string       `json:"gameId"`
	PlayerId string       `json:"playerId"`
}

func NewPlayerUpdatedIntEvent(gameIdVm string, playerIdVm string) PlayerUpdatedIntEvent {
	return PlayerUpdatedIntEvent{
		Name:     PlayerUpdatedIntEventName,
		GameId:   gameIdVm,
		PlayerId: playerIdVm,
	}
}

type UnitUpdatedIntEvent struct {
	Name   IntEventName     `json:"name"`
	GameId string           `json:"gameId"`
	Unit   viewmodel.UnitVm `json:"unit"`
}

func NewUnitUpdatedIntEvent(gameIdVm string, unitVm viewmodel.UnitVm) UnitUpdatedIntEvent {
	return UnitUpdatedIntEvent{
		Name:   UnitUpdatedIntEventName,
		GameId: gameIdVm,
		Unit:   unitVm,
	}
}
