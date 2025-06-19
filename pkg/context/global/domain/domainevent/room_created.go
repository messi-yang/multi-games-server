package domainevent

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
)

type RoomCreated struct {
	occurredOn time.Time
	roomId     globalcommonmodel.RoomId
	userId     globalcommonmodel.UserId
}

// Interface Implementation Check
var _ domain.DomainEvent = (*RoomCreated)(nil)

func NewRoomCreated(roomId globalcommonmodel.RoomId, userId globalcommonmodel.UserId) RoomCreated {
	return RoomCreated{
		occurredOn: time.Now(),
		roomId:     roomId,
		userId:     userId,
	}
}

func (domainEvent RoomCreated) GetEventName() string {
	return "ROOM_CREATED"
}

func (domainEvent RoomCreated) GetOccurredOn() time.Time {
	return domainEvent.occurredOn
}

func (domainEvent RoomCreated) GetRoomId() globalcommonmodel.RoomId {
	return domainEvent.roomId
}

func (domainEvent RoomCreated) GetUserId() globalcommonmodel.UserId {
	return domainEvent.userId
}
