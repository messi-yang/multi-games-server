package roomaccessmodel

import "github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"

type RoomMemberRepo interface {
	Add(RoomMember) error
	Get(RoomMemberId) (RoomMember, error)
	Delete(RoomMember) error
	GetRoomMemberOfUser(globalcommonmodel.RoomId, globalcommonmodel.UserId) (roomMember *RoomMember, err error)
	GetRoomMembersInRoom(globalcommonmodel.RoomId) ([]RoomMember, error)
}
