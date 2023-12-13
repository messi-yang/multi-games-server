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
	name                 string                   // The name of the player
	heldItemId           *worldcommonmodel.ItemId // Optional, The item held by the player
	action               PlayerAction
	position             worldcommonmodel.Position
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
		id:         NewPlayerId(uuid.New()),
		worldId:    worldId,
		name:       name,
		heldItemId: heldItemId,
		action: NewPlayerAction(
			PlayerActionNameEnumStand,
			worldcommonmodel.NewPosition(0, 0),
			worldcommonmodel.NewDirection(0),
			time.Now(),
		),
		position:             worldcommonmodel.NewPosition(0, 0),
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
	position worldcommonmodel.Position,
	createdAt time.Time,
	updatedAt time.Time,
) Player {
	player := Player{
		id:                   id,
		worldId:              worldId,
		userId:               userId,
		name:                 name,
		heldItemId:           heldItemId,
		action:               action,
		position:             position,
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

func (player *Player) ChangeHeldItem(itemId worldcommonmodel.ItemId) {
	player.heldItemId = &itemId
}

func (player *Player) GetHeldItemId() *worldcommonmodel.ItemId {
	return player.heldItemId
}

func (player *Player) GetAction() PlayerAction {
	return player.action
}

func (player *Player) GetPosition() worldcommonmodel.Position {
	return player.position
}

func (player *Player) Teleport(position worldcommonmodel.Position) {
	player.action = player.action.UpdatePosition(position)
}

func (player *Player) ChangeAction(action PlayerAction) {
	player.action = action
}

func (player *Player) GetCreatedAt() time.Time {
	return player.createdAt
}

func (player *Player) GetUpdatedAt() time.Time {
	return player.updatedAt
}
