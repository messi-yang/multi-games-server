package worldmodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type WorldDeleted struct {
	occurredOn time.Time
	worldId    sharedkernelmodel.WorldId
	userId     sharedkernelmodel.UserId
}

// Interface Implementation Check
var _ domain.DomainEvent = (*WorldDeleted)(nil)

func NewWorldDeleted(worldId sharedkernelmodel.WorldId, userId sharedkernelmodel.UserId) WorldDeleted {
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

func (domainEvent WorldDeleted) GetWorldId() sharedkernelmodel.WorldId {
	return domainEvent.worldId
}

func (domainEvent WorldDeleted) GetUserId() sharedkernelmodel.UserId {
	return domainEvent.userId
}
