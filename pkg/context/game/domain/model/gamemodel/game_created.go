package gamemodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
)

type GameCreated struct {
	occurredOn time.Time
	gameId     GameId
	roomId     globalcommonmodel.RoomId
}

// Interface Implementation Check
var _ domain.DomainEvent = (*GameCreated)(nil)

func NewGameCreated(gameId GameId, roomId globalcommonmodel.RoomId) GameCreated {
	return GameCreated{
		occurredOn: time.Now(),
		gameId:     gameId,
	}
}

func (domainEvent GameCreated) GetEventName() string {
	return "GAME_CREATED"
}

func (domainEvent GameCreated) GetOccurredOn() time.Time {
	return domainEvent.occurredOn
}

func (domainEvent GameCreated) GetGameId() GameId {
	return domainEvent.gameId
}

func (domainEvent GameCreated) GetRoomId() globalcommonmodel.RoomId {
	return domainEvent.roomId
}
