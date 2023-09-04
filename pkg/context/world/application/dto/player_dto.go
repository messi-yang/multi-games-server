package dto

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/playermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type PlayerDto struct {
	Id         uuid.UUID   `json:"id"`
	WorldId    uuid.UUID   `json:"worldId"`
	UserId     *uuid.UUID  `json:"userId"`
	Name       string      `json:"name"`
	Position   PositionDto `json:"position"`
	Direction  int8        `json:"direction"`
	HeldItemId *uuid.UUID  `json:"heldItemId"`
	CreatedAt  time.Time   `json:"createdAt"`
	UpdatedAt  time.Time   `json:"updatedAt"`
}

func NewPlayerDto(player playermodel.Player) PlayerDto {
	dto := PlayerDto{
		Id:      player.GetId().Uuid(),
		WorldId: player.GetWorldId().Uuid(),
		UserId: lo.TernaryF(
			player.GetUserId() == nil,
			func() *uuid.UUID { return nil },
			func() *uuid.UUID { return commonutil.ToPointer((*player.GetUserId()).Uuid()) },
		),
		Name:      player.GetName(),
		Position:  NewPositionDto(player.GetPosition()),
		Direction: player.GetDirection().Int8(),
		HeldItemId: lo.TernaryF(
			player.GetHeldItemId() == nil,
			func() *uuid.UUID { return nil },
			func() *uuid.UUID { return commonutil.ToPointer((*player.GetHeldItemId()).Uuid()) },
		),
		CreatedAt: player.GetCreatedAt(),
		UpdatedAt: player.GetCreatedAt(),
	}
	return dto
}

func ParsePlayerDto(playerDto PlayerDto) playermodel.Player {
	return playermodel.LoadPlayer(
		playermodel.NewPlayerId(playerDto.Id),
		globalcommonmodel.NewWorldId(playerDto.WorldId),
		lo.TernaryF(
			playerDto.UserId == nil,
			func() *globalcommonmodel.UserId { return nil },
			func() *globalcommonmodel.UserId {
				return commonutil.ToPointer(globalcommonmodel.NewUserId(*playerDto.UserId))
			},
		),
		playerDto.Name,
		worldcommonmodel.NewPosition(playerDto.Position.X, playerDto.Position.Z),
		worldcommonmodel.NewDirection(playerDto.Direction),
		lo.TernaryF(
			playerDto.HeldItemId == nil,
			func() *worldcommonmodel.ItemId { return nil },
			func() *worldcommonmodel.ItemId {
				return commonutil.ToPointer(worldcommonmodel.NewItemId(*playerDto.HeldItemId))
			},
		),
		playerDto.CreatedAt,
		playerDto.UpdatedAt,
	)
}
