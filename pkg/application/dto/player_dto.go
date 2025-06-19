package dto

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/gamecommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/playermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type PlayerDto struct {
	Id              uuid.UUID          `json:"id"`
	RoomId          uuid.UUID          `json:"roomId"`
	UserId          *uuid.UUID         `json:"userId"`
	Name            string             `json:"name"`
	Direction       int8               `json:"direction"`
	Action          PlayerActionDto    `json:"action"`
	PrecisePosition PrecisePositionDto `json:"precisePosition"`
	CreatedAt       time.Time          `json:"createdAt"`
	UpdatedAt       time.Time          `json:"updatedAt"`
}

func NewPlayerDto(player playermodel.Player) PlayerDto {
	dto := PlayerDto{
		Id:     player.GetId().Uuid(),
		RoomId: player.GetRoomId().Uuid(),
		UserId: lo.TernaryF(
			player.GetUserId() == nil,
			func() *uuid.UUID { return nil },
			func() *uuid.UUID { return commonutil.ToPointer((*player.GetUserId()).Uuid()) },
		),
		Name:            player.GetName(),
		Action:          NewPlayerActionDto(player.GetAction()),
		PrecisePosition: NewPrecisePositionDto(player.GetPrecisePosition()),
		CreatedAt:       player.GetCreatedAt(),
		UpdatedAt:       player.GetCreatedAt(),
	}
	return dto
}

func ParsePlayerDto(playerDto PlayerDto) (player playermodel.Player, err error) {
	action, err := ParsePlayerActionDto(playerDto.Action)
	if err != nil {
		return player, err
	}

	return playermodel.LoadPlayer(
		playermodel.NewPlayerId(playerDto.Id),
		globalcommonmodel.NewRoomId(playerDto.RoomId),
		lo.TernaryF(
			playerDto.UserId == nil,
			func() *globalcommonmodel.UserId { return nil },
			func() *globalcommonmodel.UserId {
				return commonutil.ToPointer(globalcommonmodel.NewUserId(*playerDto.UserId))
			},
		),
		playerDto.Name,
		gamecommonmodel.NewDirection(playerDto.Direction),
		action,
		gamecommonmodel.NewPrecisePosition(playerDto.PrecisePosition.X, playerDto.PrecisePosition.Z),
		playerDto.CreatedAt,
		playerDto.UpdatedAt,
	), nil
}
