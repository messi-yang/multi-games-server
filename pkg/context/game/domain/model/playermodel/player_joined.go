package playermodel

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
)

type PlayerJoined struct {
	occurredOn time.Time
	playerId   commonmodel.PlayerId
	worldId    commonmodel.WorldId
}

// Interface Implementation Check
var _ domain.DomainEvent = (*PlayerJoined)(nil)

func NewPlayerJoined(playerId commonmodel.PlayerId, worldId commonmodel.WorldId) PlayerJoined {
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

func (domainEvent PlayerJoined) GetPlayerId() commonmodel.PlayerId {
	return domainEvent.playerId
}

func (domainEvent PlayerJoined) GetWorldId() commonmodel.WorldId {
	return domainEvent.worldId
}
