package gamesocketappservice

import (
	"fmt"
)

type PlayersUpdatedIntEvent struct{}

func NewPlayersUpdatedIntEventChannel(worldIdDto string, playerIdDto string) string {
	return fmt.Sprintf("GAME_%s_PLAYER_%s_PLAYERS_UPDATED", worldIdDto, playerIdDto)
}

type UnitsUpdatedIntEvent struct{}

func NewUnitsUpdatedIntEventChannel(worldIdDto string, playerIdDto string) string {
	return fmt.Sprintf("GAME_%s_PLAYER_%s_UNITS_UPDATED", worldIdDto, playerIdDto)
}

type VisionBoundUpdatedIntEvent struct{}

func NewVisionBoundUpdatedIntEventChannel(worldIdDto string, playerIdDto string) string {
	return fmt.Sprintf("GAME_%s_PLAYER_%s_VISION_BOUND_UPDATED", worldIdDto, playerIdDto)
}
