package globalcommonmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/google/uuid"
)

type RoomId struct {
	id uuid.UUID
}

// Interface Implementation Check
var _ domain.ValueObject[RoomId] = (*RoomId)(nil)

func NewRoomId(uuid uuid.UUID) RoomId {
	return RoomId{
		id: uuid,
	}
}

func (roomId RoomId) IsEqual(otherRoomId RoomId) bool {
	return roomId.id == otherRoomId.id
}

func (roomId RoomId) Uuid() uuid.UUID {
	return roomId.id
}
