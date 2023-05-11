package playermodel

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
)

type PlayerMoved struct {
	occurredOn time.Time
	playerId   commonmodel.PlayerId
	worldId    commonmodel.WorldId
}

// Interface Implementation Check
var _ domain.DomainEvent = (*PlayerMoved)(nil)

func NewPlayerMoved(playerId commonmodel.PlayerId, worldId commonmodel.WorldId) PlayerMoved {
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

func (domainEvent PlayerMoved) GetPlayerId() commonmodel.PlayerId {
	return domainEvent.playerId
}

func (domainEvent PlayerMoved) GetWorldId() commonmodel.WorldId {
	return domainEvent.worldId
}
