package roomaccessmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
)

type RoomPermission struct {
	role globalcommonmodel.RoomRole
}

// Interface Implementation Check
var _ domain.ValueObject[RoomPermission] = (*RoomPermission)(nil)

func NewRoomPermission(role globalcommonmodel.RoomRole) RoomPermission {
	return RoomPermission{
		role: role,
	}
}

func (roomPermission RoomPermission) IsEqual(otherRoomPermission RoomPermission) bool {
	return roomPermission.role.IsEqual(otherRoomPermission.role)
}

func (roomPermission RoomPermission) CanGetRoomMembers() bool {
	return true
}

func (roomPermission RoomPermission) CanUpdateRoom() bool {
	return roomPermission.role.IsOwner() || roomPermission.role.IsAdmin()
}

func (roomPermission RoomPermission) CanDeleteRoom() bool {
	return roomPermission.role.IsOwner()
}
