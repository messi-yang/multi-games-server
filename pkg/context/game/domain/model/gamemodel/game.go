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
	selected             bool
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
	state map[string]interface{},
) Game {
	newGame := Game{
		id:                   NewGameId(uuid.New()),
		roomId:               roomId,
		name:                 name,
		started:              false,
		selected:             true,
		state:                state,
		createdAt:            time.Now(),
		updatedAt:            time.Now(),
		domainEventCollector: domain.NewDomainEventCollector(),
	}
	return newGame
}

func LoadGame(
	id GameId,
	roomId globalcommonmodel.RoomId,
	name string,
	started bool,
	selected bool,
	state map[string]interface{},
	createdAt time.Time,
	updatedAt time.Time,
) Game {
	return Game{
		id:                   id,
		roomId:               roomId,
		name:                 name,
		started:              started,
		selected:             selected,
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

func (game *Game) GetSelected() bool {
	return game.selected
}

func (game *Game) SetSelected(selected bool) {
	game.selected = selected
}

func (game *Game) GetState() map[string]interface{} {
	return game.state
}

func (game *Game) GetCreatedAt() time.Time {
	return game.createdAt
}

func (game *Game) GetUpdatedAt() time.Time {
	return game.updatedAt
}
