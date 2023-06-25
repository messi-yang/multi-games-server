package playermodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/commonmodel"
)

type Player struct {
	id                   PlayerId                  // Id of the player
	worldId              sharedkernelmodel.WorldId // The id of the world the player belongs to
	userId               *sharedkernelmodel.UserId
	name                 string                // The name of the player
	position             commonmodel.Position  // The current position of the player
	direction            commonmodel.Direction // The direction where the player is facing
	heldItemId           *commonmodel.ItemId   // Optional, The item held by the player
	createdAt            time.Time
	updatedAt            time.Time
	domainEventCollector *domain.DomainEventCollector
}

// Interface Implementation Check
var _ domain.Aggregate = (*Player)(nil)

func NewPlayer(
	id PlayerId,
	worldId sharedkernelmodel.WorldId,
	name string,
	position commonmodel.Position,
	direction commonmodel.Direction,
	heldItemId *commonmodel.ItemId,
) Player {
	player := Player{
		id:                   id,
		worldId:              worldId,
		name:                 name,
		position:             position,
		direction:            direction,
		heldItemId:           heldItemId,
		createdAt:            time.Now(),
		updatedAt:            time.Now(),
		domainEventCollector: domain.NewDomainEventCollector(),
	}
	player.domainEventCollector.Add(NewPlayerJoined(id, worldId))
	return player
}

func LoadPlayer(
	id PlayerId,
	worldId sharedkernelmodel.WorldId,
	userId *sharedkernelmodel.UserId,
	name string,
	position commonmodel.Position,
	direction commonmodel.Direction,
	heldItemId *commonmodel.ItemId,
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

func (player *Player) GetWorldId() sharedkernelmodel.WorldId {
	return player.worldId
}

func (player *Player) GetUserId() *sharedkernelmodel.UserId {
	return player.userId
}

func (player *Player) GetName() string {
	return player.name
}

func (player *Player) GetPosition() commonmodel.Position {
	return player.position
}

func (player *Player) Move(position commonmodel.Position, direction commonmodel.Direction) {
	player.position = position
	player.direction = direction
	player.domainEventCollector.Add(NewPlayerMoved(player.id, player.worldId))
}

func (player *Player) GetDirection() commonmodel.Direction {
	return player.direction
}

func (player *Player) GetPositionOneStepFoward() commonmodel.Position {
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

func (player *Player) ChangeHeldItem(itemId commonmodel.ItemId) {
	player.heldItemId = &itemId
	player.domainEventCollector.Add(NewPlayerMoved(player.id, player.worldId))
}

func (player *Player) GetHeldItemId() *commonmodel.ItemId {
	return player.heldItemId
}

func (player *Player) GetCreatedAt() time.Time {
	return player.createdAt
}

func (player *Player) GetUpdatedAt() time.Time {
	return player.updatedAt
}

func (player *Player) Delete() {
	player.domainEventCollector.Add(NewPlayerLeft(player.id, player.worldId))
}
