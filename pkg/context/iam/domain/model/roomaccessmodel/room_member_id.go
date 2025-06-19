package roomaccessmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/google/uuid"
)

type RoomMemberId struct {
	id uuid.UUID
}

// Interface Implementation Check
var _ domain.ValueObject[RoomMemberId] = (*RoomMemberId)(nil)

func NewRoomMemberId(uuid uuid.UUID) RoomMemberId {
	return RoomMemberId{
		id: uuid,
	}
}

func (itemId RoomMemberId) IsEqual(otherRoomMemberId RoomMemberId) bool {
	return itemId.id == otherRoomMemberId.id
}

func (itemId RoomMemberId) Uuid() uuid.UUID {
	return itemId.id
}
