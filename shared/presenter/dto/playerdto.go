package dto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/entity"
	"github.com/google/uuid"
)

type PlayerDto struct {
	Id uuid.UUID `json:"id"`
}

func NewPlayerDto(player entity.Player) PlayerDto {
	return PlayerDto{
		Id: player.GetId(),
	}
}

func (dto PlayerDto) ToPlayer() entity.Player {
	return entity.NewPlayerWithExistingId(dto.Id)
}
