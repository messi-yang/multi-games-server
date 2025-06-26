package gamemodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/google/uuid"
)

type Game struct {
	id                   GameId
	roomId               globalcommonmodel.RoomId
	name                 string
	started              bool
	state                map[string]interface{}
	createdAt            time.Time
	updatedAt            time.Time
	domainEventCollector *domain.DomainEventCollector
}

// Interface Implementation Check
var _ domain.DomainEventDispatchableAggregate = (*Game)(nil)

func NewGame(
	roomId globalcommonmodel.RoomId,
	name string,
) Game {
	newGame := Game{
		id:                   NewGameId(uuid.New()),
		roomId:               roomId,
		name:                 name,
		started:              false,
		state:                map[string]interface{}{},
		createdAt:            time.Now(),
		updatedAt:            time.Now(),
		domainEventCollector: domain.NewDomainEventCollector(),
	}
	newGame.domainEventCollector.Add(NewGameCreated(
		newGame.id,
		newGame.roomId,
	))
	return newGame
}

func LoadGame(
	id GameId,
	roomId globalcommonmodel.RoomId,
	name string,
	started bool,
	state map[string]interface{},
	createdAt time.Time,
	updatedAt time.Time,
) Game {
	return Game{
		id:                   id,
		roomId:               roomId,
		name:                 name,
		started:              started,
		state:                state,
		createdAt:            createdAt,
		updatedAt:            updatedAt,
		domainEventCollector: domain.NewDomainEventCollector(),
	}
}

func (game *Game) PopDomainEvents() []domain.DomainEvent {
	return game.domainEventCollector.PopAll()
}

func (game *Game) GetId() GameId {
	return game.id
}

func (game *Game) GetRoomId() globalcommonmodel.RoomId {
	return game.roomId
}

func (game *Game) GetName() string {
	return game.name
}

func (game *Game) GetStarted() bool {
	return game.started
}

func (game *Game) SetStarted(started bool) {
	game.started = started
}

func (game *Game) GetState() map[string]interface{} {
	return game.state
}

func (game *Game) SetState(state map[string]interface{}) {
	game.state = state
}

func (game *Game) GetCreatedAt() time.Time {
	return game.createdAt
}

func (game *Game) GetUpdatedAt() time.Time {
	return game.updatedAt
}
