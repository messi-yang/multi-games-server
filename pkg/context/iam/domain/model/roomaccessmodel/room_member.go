package roomaccessmodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/google/uuid"
)

type RoomMember struct {
	id        RoomMemberId
	roomId    globalcommonmodel.RoomId
	userId    globalcommonmodel.UserId
	role      globalcommonmodel.RoomRole
	createdAt time.Time
	updatedAt time.Time
}

// Interface Implementation Check
var _ domain.Aggregate = (*RoomMember)(nil)

func NewRoomMember(
	roomId globalcommonmodel.RoomId,
	userId globalcommonmodel.UserId,
	role globalcommonmodel.RoomRole,
) RoomMember {
	newRoomRole := RoomMember{
		id:        NewRoomMemberId(uuid.New()),
		roomId:    roomId,
		userId:    userId,
		role:      role,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
	return newRoomRole
}

func LoadRoomMember(
	id RoomMemberId,
	roomId globalcommonmodel.RoomId,
	userId globalcommonmodel.UserId,
	role globalcommonmodel.RoomRole,
	createdAt time.Time,
	updatedAt time.Time,
) RoomMember {
	return RoomMember{
		id:        id,
		roomId:    roomId,
		userId:    userId,
		role:      role,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (roomMember *RoomMember) GetId() RoomMemberId {
	return roomMember.id
}

func (roomMember *RoomMember) GeRoomId() globalcommonmodel.RoomId {
	return roomMember.roomId
}

func (roomMember *RoomMember) GeUserId() globalcommonmodel.UserId {
	return roomMember.userId
}

func (roomMember *RoomMember) GetRole() globalcommonmodel.RoomRole {
	return roomMember.role
}

func (roomMember *RoomMember) GetCreatedAt() time.Time {
	return roomMember.createdAt
}

func (roomMember *RoomMember) GetUpdatedAt() time.Time {
	return roomMember.updatedAt
}
