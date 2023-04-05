package jsondto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/playermodel"
	"github.com/google/uuid"
)

type PlayerAggDto struct {
	Id          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Position    PositionVoDto `json:"position"`
	Direction   int8          `json:"direction"`
	VisionBound BoundVoDto    `json:"visionBound"`
	HeldItemId  *uuid.UUID    `json:"heldItemId"`
}

func NewPlayerAggDto(player playermodel.PlayerAgg) PlayerAggDto {
	dto := PlayerAggDto{
		Id:          player.GetId().Uuid(),
		Name:        player.GetName(),
		Position:    NewPositionVoDto(player.GetPosition()),
		Direction:   player.GetDirection().Int8(),
		VisionBound: NewBoundVoDto(player.GetVisionBound()),
	}
	playerHeldItemid := player.GetHeldItemId()
	if playerHeldItemid == nil {
		dto.HeldItemId = nil
	} else {
		heldItemIdDto := (*player.GetHeldItemId()).Uuid()
		dto.HeldItemId = &heldItemIdDto
	}
	return dto
}
