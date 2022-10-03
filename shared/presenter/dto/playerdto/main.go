package playerdto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/entity"
	"github.com/google/uuid"
)

type Dto struct {
	Id uuid.UUID `json:"id"`
}

func ToDto(player entity.Player) Dto {
	return Dto{
		Id: player.GetId(),
	}
}

func FromDto(player Dto) entity.Player {
	return entity.NewPlayerWithExistingId(player.Id)
}
