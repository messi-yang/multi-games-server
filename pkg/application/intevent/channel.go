package intevent

import (
	"fmt"
)

func CreateGameClientChannel(gameId string, playerId string) string {
	return fmt.Sprintf("GAME_%s_CLIENT_PLAYER_%s", gameId, playerId)
}
