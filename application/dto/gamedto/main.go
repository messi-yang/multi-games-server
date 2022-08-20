package gamedto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/mapsizedto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/unitdto"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/entity"
	"github.com/google/uuid"
)

type GameDTO struct {
	Id      uuid.UUID
	UnitMap [][]unitdto.UnitDTO
	MapSize mapsizedto.MapSizeDTO
}

func ToDTO(game entity.Game) GameDTO {
	return GameDTO{
		Id:      game.GetId(),
		UnitMap: unitdto.ToDTOMap(game.GetUnitMap()),
		MapSize: mapsizedto.ToDTO(game.GetMapSize()),
	}
}
