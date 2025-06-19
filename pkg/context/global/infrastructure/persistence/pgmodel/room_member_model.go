package pgmodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/roomaccessmodel"
	"github.com/google/uuid"
)

type RoomRole string

const (
	RoomRoleOwner  RoomRole = "owner"
	RoomRoleAdmin  RoomRole = "admin"
	RoomRoleEditor RoomRole = "editor"
	RoomRoleViewer RoomRole = "viewer"
)

type RoomMemberModel struct {
	Id        uuid.UUID `gorm:"primaryKey"`
	UserId    uuid.UUID `gorm:"not null"`
	RoomId    uuid.UUID `gorm:"not null"`
	Role      RoomRole  `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;not null"`
}

func (RoomMemberModel) TableName() string {
	return "room_members"
}

func NewRoomMemberModel(roomMember roomaccessmodel.RoomMember) RoomMemberModel {
	return RoomMemberModel{
		Id:        roomMember.GetId().Uuid(),
		RoomId:    roomMember.GeRoomId().Uuid(),
		UserId:    roomMember.GeUserId().Uuid(),
		Role:      RoomRole(roomMember.GetRole().String()),
		CreatedAt: roomMember.GetCreatedAt(),
		UpdatedAt: roomMember.GetUpdatedAt(),
	}
}

func ParseRoomMemberModel(roomMemberModel RoomMemberModel) (roomMember roomaccessmodel.RoomMember, err error) {
	roomRole, err := globalcommonmodel.NewRoomRole(string(roomMemberModel.Role))
	if err != nil {
		return roomMember, err
	}
	return roomaccessmodel.LoadRoomMember(
		roomaccessmodel.NewRoomMemberId(roomMemberModel.Id),
		globalcommonmodel.NewRoomId(roomMemberModel.RoomId),
		globalcommonmodel.NewUserId(roomMemberModel.UserId),
		roomRole,
		roomMemberModel.CreatedAt,
		roomMemberModel.UpdatedAt,
	), nil
}
