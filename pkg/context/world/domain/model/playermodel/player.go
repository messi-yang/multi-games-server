package playermodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/google/uuid"
)

type Player struct {
	id                   PlayerId                  // Id of the player
	worldId              globalcommonmodel.WorldId // The id of the world the player belongs to
	userId               *globalcommonmodel.UserId
	name                 string                     // The name of the player
	direction            worldcommonmodel.Direction // The direction where the player is facing
	heldItemId           *worldcommonmodel.ItemId   // Optional, The item held by the player
	action               PlayerAction
	actionPosition       worldcommonmodel.Position
	actedAt              time.Time
	createdAt            time.Time
	updatedAt            time.Time
	domainEventCollector *domain.DomainEventCollector
}

// Interface Implementation Check
var _ domain.Aggregate = (*Player)(nil)

func NewPlayer(
	worldId globalcommonmodel.WorldId,
	name string,
	direction worldcommonmodel.Direction,
	heldItemId *worldcommonmodel.ItemId,
) Player {
	return Player{
		id:                   NewPlayerId(uuid.New()),
		worldId:              worldId,
		name:                 name,
		direction:            direction,
		heldItemId:           heldItemId,
		action:               NewPlayerActionStand(),
		actionPosition:       worldcommonmodel.NewPosition(0, 0),
		actedAt:              time.Now(),
		createdAt:            time.Now(),
		updatedAt:            time.Now(),
		domainEventCollector: domain.NewDomainEventCollector(),
	}
}

func LoadPlayer(
	id PlayerId,
	worldId globalcommonmodel.WorldId,
	userId *globalcommonmodel.UserId,
	name string,
	direction worldcommonmodel.Direction,
	heldItemId *worldcommonmodel.ItemId,
	action PlayerAction,
	actionPosition worldcommonmodel.Position,
	actedAt time.Time,
	createdAt time.Time,
	updatedAt time.Time,
) Player {
	player := Player{
		id:                   id,
		worldId:              worldId,
		userId:               userId,
		name:                 name,
		direction:            direction,
		heldItemId:           heldItemId,
		action:               action,
		actionPosition:       actionPosition,
		actedAt:              actedAt,
		createdAt:            createdAt,
		updatedAt:            updatedAt,
		domainEventCollector: domain.NewDomainEventCollector(),
	}
	return player
}

func (player *Player) PopDomainEvents() []domain.DomainEvent {
	return player.domainEventCollector.PopAll()
}

func (player *Player) GetId() PlayerId {
	return player.id
}

func (player *Player) GetWorldId() globalcommonmodel.WorldId {
	return player.worldId
}

func (player *Player) GetUserId() *globalcommonmodel.UserId {
	return player.userId
}

func (player *Player) GetName() string {
	return player.name
}

func (player *Player) Move(position worldcommonmodel.Position, direction worldcommonmodel.Direction) {
	// player.position = position
	player.direction = direction
}

func (player *Player) GetDirection() worldcommonmodel.Direction {
	return player.direction
}

func (player *Player) ChangeHeldItem(itemId worldcommonmodel.ItemId) {
	player.heldItemId = &itemId
}

func (player *Player) GetHeldItemId() *worldcommonmodel.ItemId {
	return player.heldItemId
}

func (player *Player) GetAction() PlayerAction {
	return player.action
}

func (player *Player) GetActionPosition() worldcommonmodel.Position {
	return player.actionPosition
}

func (player *Player) Teleport(position worldcommonmodel.Position) {
	player.actionPosition = position
	player.actedAt = time.Now()
}

func (player *Player) Walk(actionPosition worldcommonmodel.Position, direction worldcommonmodel.Direction) {
	player.action = NewPlayerActionWalk()
	player.actedAt = time.Now()
	player.actionPosition = actionPosition
	player.direction = direction
}

func (player *Player) Stand(actionPosition worldcommonmodel.Position, direction worldcommonmodel.Direction) {
	player.action = NewPlayerActionWalk()
	player.actedAt = time.Now()
	player.actionPosition = actionPosition
	player.direction = direction
}

func (player *Player) GetActedAt() time.Time {
	return player.actedAt
}

func (player *Player) GetCreatedAt() time.Time {
	return player.createdAt
}

func (player *Player) GetUpdatedAt() time.Time {
	return player.updatedAt
}
