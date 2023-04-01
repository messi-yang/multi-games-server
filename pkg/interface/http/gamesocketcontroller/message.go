package gamesocketcontroller

import (
	"fmt"

	"github.com/google/uuid"
)

type PlayersUpdatedMessage struct{}

func newPlayersUpdatedMessageChannel(worldIdDto uuid.UUID) string {
	return fmt.Sprintf("GAME_%s_PLAYERS_UPDATED", worldIdDto)
}

type UnitsUpdatedMessage struct{}

func NewUnitsUpdatedMessageChannel(worldIdDto uuid.UUID) string {
	return fmt.Sprintf("GAME_%s_UNITS_UPDATED", worldIdDto)
}
