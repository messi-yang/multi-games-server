package pgmodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/roommodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/google/uuid"
)

type RoomModel struct {
	Id        uuid.UUID `gorm:"primaryKey"`
	UserId    uuid.UUID `gorm:"unique;not null"`
	Name      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;not null"`
}

func (RoomModel) TableName() string {
	return "rooms"
}

func NewRoomModel(room roommodel.Room) RoomModel {
	return RoomModel{
		Id:        room.GetId().Uuid(),
		UserId:    room.GetUserId().Uuid(),
		Name:      room.GetName(),
		UpdatedAt: room.GetUpdatedAt(),
		CreatedAt: room.GetCreatedAt(),
	}
}

func ParseRoomModel(roomModel RoomModel) (room roommodel.Room) {
	return roommodel.LoadRoom(
		globalcommonmodel.NewRoomId(roomModel.Id),
		globalcommonmodel.NewUserId(roomModel.UserId),
		roomModel.Name,
		roomModel.CreatedAt,
		roomModel.UpdatedAt,
	)
}
