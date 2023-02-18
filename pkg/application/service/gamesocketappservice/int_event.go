package gamesocketappservice

import (
	"fmt"
)

type PlayersUpdatedIntEvent struct{}

func NewPlayersUpdatedIntEventChannel(gameIdDto string, playerIdDto string) string {
	return fmt.Sprintf("GAME_%s_PLAYER_%s_PLAYERS_UPDATED", gameIdDto, playerIdDto)
}

type UnitsUpdatedIntEvent struct{}

func NewUnitsUpdatedIntEventChannel(gameIdDto string, playerIdDto string) string {
	return fmt.Sprintf("GAME_%s_PLAYER_%s_UNITS_UPDATED", gameIdDto, playerIdDto)
}
