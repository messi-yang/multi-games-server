package pgmodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/gamemodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/roommodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"github.com/google/uuid"
)

type RoomModel struct {
	Id            uuid.UUID  `gorm:"primaryKey"`
	UserId        uuid.UUID  `gorm:"unique;not null"`
	Name          string     `gorm:"not null"`
	CurrentGameId *uuid.UUID `gorm:""`
	CreatedAt     time.Time  `gorm:"autoCreateTime;not null"`
	UpdatedAt     time.Time  `gorm:"autoUpdateTime;not null"`
}

func (RoomModel) TableName() string {
	return "rooms"
}

func NewRoomModel(room roommodel.Room) RoomModel {
	var currentGameId *uuid.UUID
	if room.GetCurrentGameId() != nil {
		currentGameId = commonutil.ToPointer(room.GetCurrentGameId().Uuid())
	}

	return RoomModel{
		Id:            room.GetId().Uuid(),
		UserId:        room.GetUserId().Uuid(),
		Name:          room.GetName(),
		CurrentGameId: currentGameId,
		UpdatedAt:     room.GetUpdatedAt(),
		CreatedAt:     room.GetCreatedAt(),
	}
}

func ParseRoomModel(roomModel RoomModel) (room roommodel.Room) {
	var currentGameId *gamemodel.GameId
	if roomModel.CurrentGameId != nil {
		currentGameId = commonutil.ToPointer(gamemodel.NewGameId(*roomModel.CurrentGameId))
	}

	return roommodel.LoadRoom(
		globalcommonmodel.NewRoomId(roomModel.Id),
		globalcommonmodel.NewUserId(roomModel.UserId),
		roomModel.Name,
		currentGameId,
		roomModel.CreatedAt,
		roomModel.UpdatedAt,
	)
}
