package roommodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
)

type RoomDeleted struct {
	occurredOn time.Time
	roomId     globalcommonmodel.RoomId
	userId     globalcommonmodel.UserId
}

// Interface Implementation Check
var _ domain.DomainEvent = (*RoomDeleted)(nil)

func NewRoomDeleted(roomId globalcommonmodel.RoomId, userId globalcommonmodel.UserId) RoomDeleted {
	return RoomDeleted{
		occurredOn: time.Now(),
		roomId:     roomId,
		userId:     userId,
	}
}

func (domainEvent RoomDeleted) GetEventName() string {
	return "ROOM_DELETED"
}

func (domainEvent RoomDeleted) GetOccurredOn() time.Time {
	return domainEvent.occurredOn
}

func (domainEvent RoomDeleted) GetRoomId() globalcommonmodel.RoomId {
	return domainEvent.roomId
}

func (domainEvent RoomDeleted) GetUserId() globalcommonmodel.UserId {
	return domainEvent.userId
}
