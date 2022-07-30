package gameroomdto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/gamedto"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/aggregate"
)

type GameRoomDTO struct {
	Game gamedto.GameDTO
}

func ToDTO(gameRoom aggregate.GameRoom) GameRoomDTO {
	return GameRoomDTO{
		Game: gamedto.ToDTO(gameRoom.GetGame()),
	}
}
