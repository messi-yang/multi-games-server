package dto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/playermodel"
	"github.com/google/uuid"
)

type PlayerDto struct {
	Id          uuid.UUID   `json:"id"`
	Name        string      `json:"name"`
	Position    PositionDto `json:"position"`
	Direction   int8        `json:"direction"`
	VisionBound BoundDto    `json:"visionBound"`
	HeldItemId  *uuid.UUID  `json:"heldItemId"`
}

func NewPlayerDto(player playermodel.Player) PlayerDto {
	dto := PlayerDto{
		Id:          player.GetId().Uuid(),
		Name:        player.GetName(),
		Position:    NewPositionDto(player.GetPosition()),
		Direction:   player.GetDirection().Int8(),
		VisionBound: NewBoundDto(player.GetVisionBound()),
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
