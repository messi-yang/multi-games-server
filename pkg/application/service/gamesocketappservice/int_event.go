package gamesocketappservice

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/json"
)

type PlayersUpdatedIntEvent struct{}

func NewPlayersUpdatedIntEventChannel(gameIdDto string, playerIdDto string) string {
	return fmt.Sprintf("GAME_%s_PLAYER_%s_PLAYERS_UPDATED", gameIdDto, playerIdDto)
}
func (event PlayersUpdatedIntEvent) Marshal() []byte {
	return json.Marshal(event)
}

type ViewUpdatedIntEvent struct{}

func NewViewUpdatedIntEventChannel(gameIdDto string, playerIdDto string) string {
	return fmt.Sprintf("GAME_%s_PLAYER_%s_VIEW_UPDATED", gameIdDto, playerIdDto)
}
func (event ViewUpdatedIntEvent) Marshal() []byte {
	return json.Marshal(event)
}
