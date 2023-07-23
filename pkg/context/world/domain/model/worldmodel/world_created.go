package worldmodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
)

type WorldCreated struct {
	occurredOn time.Time
	worldId    globalcommonmodel.WorldId
	userId     globalcommonmodel.UserId
}

// Interface Implementation Check
var _ domain.DomainEvent = (*WorldCreated)(nil)

func NewWorldCreated(worldId globalcommonmodel.WorldId, userId globalcommonmodel.UserId) WorldCreated {
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

func (domainEvent WorldCreated) GetWorldId() globalcommonmodel.WorldId {
	return domainEvent.worldId
}

func (domainEvent WorldCreated) GetUserId() globalcommonmodel.UserId {
	return domainEvent.userId
}
