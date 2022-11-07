package dto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game/valueobject"
	"github.com/google/uuid"
)

type PlayerIdDto struct {
	Value uuid.UUID `json:"value"`
}

func NewPlayerIdDto(id uuid.UUID) PlayerIdDto {
	return PlayerIdDto{
		Value: id,
	}
}

func (dto PlayerIdDto) ToValueObject() valueobject.PlayerId {
	return valueobject.NewPlayerId(dto.Value)
}
