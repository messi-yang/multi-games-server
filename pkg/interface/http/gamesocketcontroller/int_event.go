package gamesocketcontroller

import (
	"fmt"

	"github.com/google/uuid"
)

type PlayersUpdatedIntEvent struct{}

func newPlayersUpdatedIntEventChannel(worldIdDto uuid.UUID) string {
	return fmt.Sprintf("GAME_%s_PLAYERS_UPDATED", worldIdDto)
}

type UnitsUpdatedIntEvent struct{}

func NewUnitsUpdatedIntEventChannel(worldIdDto uuid.UUID) string {
	return fmt.Sprintf("GAME_%s_UNITS_UPDATED", worldIdDto)
}
