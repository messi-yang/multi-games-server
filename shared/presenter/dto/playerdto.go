package dto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/aggregate"
	"github.com/google/uuid"
)

type PlayerDto struct {
	Id uuid.UUID `json:"id"`
}

func NewPlayerDto(player aggregate.Player) PlayerDto {
	return PlayerDto{
		Id: player.GetId(),
	}
}

func (dto PlayerDto) ToEntity() aggregate.Player {
	return aggregate.NewPlayer(dto.Id)
}
