package globalcommonmodel

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
)

type RoomRoleValue string

const (
	RoomRoleOwner  RoomRoleValue = "owner"
	RoomRoleAdmin  RoomRoleValue = "admin"
	RoomRoleEditor RoomRoleValue = "editor"
	RoomRoleViewer RoomRoleValue = "viewer"
)

type RoomRole struct {
	value RoomRoleValue
}

// Interface Implementation Check
var _ domain.ValueObject[RoomRole] = (*RoomRole)(nil)

func NewRoomRole(roomRoleValue string) (RoomRole, error) {
	switch roomRoleValue {
	case "owner":
		return RoomRole{
			value: RoomRoleOwner,
		}, nil
	case "admin":
		return RoomRole{
			value: RoomRoleAdmin,
		}, nil
	case "editor":
		return RoomRole{
			value: RoomRoleEditor,
		}, nil
	case "viewer":
		return RoomRole{
			value: RoomRoleViewer,
		}, nil
	default:
		return RoomRole{}, fmt.Errorf("invalid RoomRole: %s", roomRoleValue)
	}
}

func (roomRole RoomRole) IsEqual(otherRoomRole RoomRole) bool {
	return roomRole.value == otherRoomRole.value
}

func (roomRole RoomRole) String() string {
	return string(roomRole.value)
}

func (roomRole RoomRole) IsOwner() bool {
	return roomRole.value == RoomRoleOwner
}

func (roomRole RoomRole) IsAdmin() bool {
	return roomRole.value == RoomRoleAdmin
}

func (roomRole RoomRole) IsEditor() bool {
	return roomRole.value == RoomRoleEditor
}

func (roomRole RoomRole) IsViewer() bool {
	return roomRole.value == RoomRoleViewer
}
