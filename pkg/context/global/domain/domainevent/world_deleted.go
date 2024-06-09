package domainevent

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
)

type WorldDeleted struct {
	occurredOn time.Time
	worldId    globalcommonmodel.WorldId
	userId     globalcommonmodel.UserId
}

// Interface Implementation Check
var _ domain.DomainEvent = (*WorldDeleted)(nil)

func NewWorldDeleted(worldId globalcommonmodel.WorldId, userId globalcommonmodel.UserId) WorldDeleted {
	return WorldDeleted{
		occurredOn: time.Now(),
		worldId:    worldId,
		userId:     userId,
	}
}

func (domainEvent WorldDeleted) GetEventName() string {
	return "WORLD_DELETED"
}

func (domainEvent WorldDeleted) GetOccurredOn() time.Time {
	return domainEvent.occurredOn
}

func (domainEvent WorldDeleted) GetWorldId() globalcommonmodel.WorldId {
	return domainEvent.worldId
}

func (domainEvent WorldDeleted) GetUserId() globalcommonmodel.UserId {
	return domainEvent.userId
}
