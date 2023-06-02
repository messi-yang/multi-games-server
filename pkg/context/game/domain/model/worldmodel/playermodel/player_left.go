package playermodel

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type PlayerLeft struct {
	occurredOn time.Time
	playerId   PlayerId
	worldId    sharedkernelmodel.WorldId
}

// Interface Implementation Check
var _ domain.DomainEvent = (*PlayerLeft)(nil)

func NewPlayerLeft(playerId PlayerId, worldId sharedkernelmodel.WorldId) PlayerLeft {
	return PlayerLeft{
		occurredOn: time.Now(),
		playerId:   playerId,
		worldId:    worldId,
	}
}

func (domainEvent PlayerLeft) GetEventName() string {
	return "PLAYER_LEFT"
}

func (domainEvent PlayerLeft) GetOccurredOn() time.Time {
	return domainEvent.occurredOn
}

func (domainEvent PlayerLeft) GetPlayerId() PlayerId {
	return domainEvent.playerId
}

func (domainEvent PlayerLeft) GetWorldId() sharedkernelmodel.WorldId {
	return domainEvent.worldId
}
