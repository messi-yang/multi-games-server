package sharedkernelmodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain"
)

type WorldCreated struct {
	occurredOn time.Time
	worldId    WorldId
	userId     UserId
}

// Interface Implementation Check
var _ domain.DomainEvent = (*WorldCreated)(nil)

func NewWorldCreated(worldId WorldId, userId UserId) WorldCreated {
	return WorldCreated{
		occurredOn: time.Now(),
		worldId:    worldId,
		userId:     userId,
	}
}

func (domainEvent WorldCreated) GetEventName() string {
	return "WORLD_CREATED"
}

func (domainEvent WorldCreated) GetOccurredOn() time.Time {
	return domainEvent.occurredOn
}

func (domainEvent WorldCreated) GetWorldId() WorldId {
	return domainEvent.worldId
}

func (domainEvent WorldCreated) GetUserId() UserId {
	return domainEvent.userId
}
