package worldmodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type WorldCreated struct {
	occurredOn time.Time
	worldId    sharedkernelmodel.WorldId
	userId     sharedkernelmodel.UserId
}

// Interface Implementation Check
var _ domain.DomainEvent = (*WorldCreated)(nil)

func NewWorldCreated(worldId sharedkernelmodel.WorldId, userId sharedkernelmodel.UserId) WorldCreated {
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

func (domainEvent WorldCreated) GetWorldId() sharedkernelmodel.WorldId {
	return domainEvent.worldId
}

func (domainEvent WorldCreated) GetUserId() sharedkernelmodel.UserId {
	return domainEvent.userId
}
