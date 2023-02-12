package gamesocketservice

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/viewmodel"
)

func CreateGameIntEventChannel(gameId string) string {
	return fmt.Sprintf("GAME_%s", gameId)
}

type GameSocketIntEventName string

const (
	PlayerUpdatedGameSocketIntEventName GameSocketIntEventName = "PLAYER_UPDATED"
	UnitUpdatedGameSocketIntEventName   GameSocketIntEventName = "UNIT_UPDATED"
)

type GameSocketIntEvent struct {
	Name GameSocketIntEventName `json:"name"`
}

type PlayerUpdatedIntEvent struct {
	Name     GameSocketIntEventName `json:"name"`
	GameId   string                 `json:"gameId"`
	PlayerId string                 `json:"playerId"`
}

func NewPlayerUpdatedIntEvent(gameIdVm string, playerIdVm string) PlayerUpdatedIntEvent {
	return PlayerUpdatedIntEvent{
		Name:     PlayerUpdatedGameSocketIntEventName,
		GameId:   gameIdVm,
		PlayerId: playerIdVm,
	}
}

type UnitUpdatedIntEvent struct {
	Name     GameSocketIntEventName `json:"name"`
	GameId   string                 `json:"gameId"`
	Location viewmodel.LocationVm   `json:"location"`
}

func NewUnitUpdatedIntEvent(gameIdVm string, locationVm viewmodel.LocationVm) UnitUpdatedIntEvent {
	return UnitUpdatedIntEvent{
		Name:     UnitUpdatedGameSocketIntEventName,
		GameId:   gameIdVm,
		Location: locationVm,
	}
}
