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
	position             worldcommonmodel.Position  // The current position of the player
	direction            worldcommonmodel.Direction // The direction where the player is facing
	heldItemId           *worldcommonmodel.ItemId   // Optional, The item held by the player
	createdAt            time.Time
	updatedAt            time.Time
	domainEventCollector *domain.DomainEventCollector
}

// Interface Implementation Check
var _ domain.Aggregate = (*Player)(nil)

func NewPlayer(
	worldId globalcommonmodel.WorldId,
	name string,
	position worldcommonmodel.Position,
	direction worldcommonmodel.Direction,
	heldItemId *worldcommonmodel.ItemId,
) Player {
	return Player{
		id:                   NewPlayerId(uuid.New()),
		worldId:              worldId,
		name:                 name,
		position:             position,
		direction:            direction,
		heldItemId:           heldItemId,
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
	position worldcommonmodel.Position,
	direction worldcommonmodel.Direction,
	heldItemId *worldcommonmodel.ItemId,
	createdAt time.Time,
	updatedAt time.Time,
) Player {
	player := Player{
		id:                   id,
		worldId:              worldId,
		userId:               userId,
		name:                 name,
		position:             position,
		direction:            direction,
		heldItemId:           heldItemId,
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

func (player *Player) GetPosition() worldcommonmodel.Position {
	return player.position
}

func (player *Player) Move(position worldcommonmodel.Position, direction worldcommonmodel.Direction) {
	player.position = position
	player.direction = direction
}

func (player *Player) GetDirection() worldcommonmodel.Direction {
	return player.direction
}

func (player *Player) GetPositionOneStepFoward() worldcommonmodel.Position {
	direction := player.direction
	position := player.position

	if direction.IsUp() {
		return position.Shift(0, -1)
	} else if direction.IsRight() {
		return position.Shift(1, 0)
	} else if direction.IsDown() {
		return position.Shift(0, 1)
	} else if direction.IsLeft() {
		return position.Shift(-1, 0)
	} else {
		return position.Shift(0, 1)
	}
}

func (player *Player) ChangeHeldItem(itemId worldcommonmodel.ItemId) {
	player.heldItemId = &itemId
}

func (player *Player) GetHeldItemId() *worldcommonmodel.ItemId {
	return player.heldItemId
}

func (player *Player) GetCreatedAt() time.Time {
	return player.createdAt
}

func (player *Player) GetUpdatedAt() time.Time {
	return player.updatedAt
}
