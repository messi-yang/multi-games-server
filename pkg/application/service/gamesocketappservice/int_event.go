package gamesocketappservice

import (
	"fmt"

	"github.com/google/uuid"
)

type PlayersUpdatedIntEvent struct{}

func NewPlayersUpdatedIntEventChannel(worldIdDto uuid.UUID, playerIdDto uuid.UUID) string {
	return fmt.Sprintf("GAME_%s_PLAYER_%s_PLAYERS_UPDATED", worldIdDto, playerIdDto)
}

type UnitsUpdatedIntEvent struct{}

func NewUnitsUpdatedIntEventChannel(worldIdDto uuid.UUID, playerIdDto uuid.UUID) string {
	return fmt.Sprintf("GAME_%s_PLAYER_%s_UNITS_UPDATED", worldIdDto, playerIdDto)
}

type VisionBoundUpdatedIntEvent struct{}

func NewVisionBoundUpdatedIntEventChannel(worldIdDto uuid.UUID, playerIdDto uuid.UUID) string {
	return fmt.Sprintf("GAME_%s_PLAYER_%s_VISION_BOUND_UPDATED", worldIdDto, playerIdDto)
}
