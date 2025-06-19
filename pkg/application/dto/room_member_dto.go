package dto

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/roomaccessmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/usermodel"
	"github.com/google/uuid"
)

type RoomMemberDto struct {
	Id        uuid.UUID `json:"id"`
	User      UserDto   `json:"user"`
	RoomId    uuid.UUID `json:"roomId"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewRoomMemberDto(roomMember roomaccessmodel.RoomMember, user usermodel.User) RoomMemberDto {
	dto := RoomMemberDto{
		Id:        roomMember.GetId().Uuid(),
		User:      NewUserDto(user),
		RoomId:    roomMember.GeRoomId().Uuid(),
		Role:      roomMember.GetRole().String(),
		CreatedAt: roomMember.GetCreatedAt(),
		UpdatedAt: roomMember.GetUpdatedAt(),
	}
	return dto
}
