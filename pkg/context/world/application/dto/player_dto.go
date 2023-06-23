package dto

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel/playermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type PlayerDto struct {
	Id          uuid.UUID   `json:"id"`
	UserId      *uuid.UUID  `json:"userId"`
	Name        string      `json:"name"`
	Position    PositionDto `json:"position"`
	Direction   int8        `json:"direction"`
	VisionBound BoundDto    `json:"visionBound"`
	HeldItemId  *uuid.UUID  `json:"heldItemId"`
}

func NewPlayerDto(player playermodel.Player) PlayerDto {
	dto := PlayerDto{
		Id: player.GetId().Uuid(),
		UserId: lo.TernaryF(
			player.GetUserId() == nil,
			func() *uuid.UUID { return nil },
			func() *uuid.UUID { return commonutil.ToPointer((*player.GetUserId()).Uuid()) },
		),
		Name:        player.GetName(),
		Position:    NewPositionDto(player.GetPosition()),
		Direction:   player.GetDirection().Int8(),
		VisionBound: NewBoundDto(player.GetVisionBound()),
		HeldItemId: lo.TernaryF(
			player.GetHeldItemId() == nil,
			func() *uuid.UUID { return nil },
			func() *uuid.UUID { return commonutil.ToPointer((*player.GetHeldItemId()).Uuid()) },
		),
	}
	return dto
}
