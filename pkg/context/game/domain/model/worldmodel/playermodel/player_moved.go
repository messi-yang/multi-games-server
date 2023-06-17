package playermodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type PlayerMoved struct {
	occurredOn time.Time
	playerId   PlayerId
	worldId    sharedkernelmodel.WorldId
}

// Interface Implementation Check
var _ domain.DomainEvent = (*PlayerMoved)(nil)

func NewPlayerMoved(playerId PlayerId, worldId sharedkernelmodel.WorldId) PlayerMoved {
	return PlayerMoved{
		occurredOn: time.Now(),
		playerId:   playerId,
		worldId:    worldId,
	}
}

func (domainEvent PlayerMoved) GetEventName() string {
	return "PLAYER_MOVED"
}

func (domainEvent PlayerMoved) GetOccurredOn() time.Time {
	return domainEvent.occurredOn
}

func (domainEvent PlayerMoved) GetPlayerId() PlayerId {
	return domainEvent.playerId
}

func (domainEvent PlayerMoved) GetWorldId() sharedkernelmodel.WorldId {
	return domainEvent.worldId
}
