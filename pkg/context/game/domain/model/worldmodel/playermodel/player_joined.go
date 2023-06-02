package playermodel

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type PlayerJoined struct {
	occurredOn time.Time
	playerId   PlayerId
	worldId    sharedkernelmodel.WorldId
}

// Interface Implementation Check
var _ domain.DomainEvent = (*PlayerJoined)(nil)

func NewPlayerJoined(playerId PlayerId, worldId sharedkernelmodel.WorldId) PlayerJoined {
	return PlayerJoined{
		occurredOn: time.Now(),
		playerId:   playerId,
		worldId:    worldId,
	}
}

func (domainEvent PlayerJoined) GetEventName() string {
	return "PLAYER_JOINED"
}

func (domainEvent PlayerJoined) GetOccurredOn() time.Time {
	return domainEvent.occurredOn
}

func (domainEvent PlayerJoined) GetPlayerId() PlayerId {
	return domainEvent.playerId
}

func (domainEvent PlayerJoined) GetWorldId() sharedkernelmodel.WorldId {
	return domainEvent.worldId
}
