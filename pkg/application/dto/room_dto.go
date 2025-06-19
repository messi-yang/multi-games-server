package dto

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/roommodel"
	"github.com/google/uuid"
)

type RoomDto struct {
	Id        uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"userId"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewRoomDto(room roommodel.Room) RoomDto {
	return RoomDto{
		Id:        room.GetId().Uuid(),
		UserId:    room.GetUserId().Uuid(),
		Name:      room.GetName(),
		CreatedAt: room.GetCreatedAt(),
		UpdatedAt: room.GetUpdatedAt(),
	}
}
