package gamedto

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/dto/mapsizedto"
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/dto/unitdto"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/entity"
	"github.com/google/uuid"
)

type GameDTO struct {
	Id         uuid.UUID
	UnitMatrix [][]unitdto.UnitDTO
	MapSize    mapsizedto.MapSizeDTO
}

func ToDTO(game entity.Game) GameDTO {
	return GameDTO{
		Id:         game.Id,
		UnitMatrix: unitdto.ToDTOMatrix(game.UnitMatrix),
		MapSize:    mapsizedto.ToDTO(game.MapSize),
	}
}
