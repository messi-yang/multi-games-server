package intevent

import (
	"fmt"
)

func CreateGameAdminChannel() string {
	return "GAME_ADMIN"
}

func CreateGameClientChannel(gameId string, playerId string) string {
	return fmt.Sprintf("GAME_%s_CLIENT_PLAYER_%s", gameId, playerId)
}
