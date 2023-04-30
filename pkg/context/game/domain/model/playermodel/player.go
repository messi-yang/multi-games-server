package playermodel

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/domainmodel"
)

func calculatePlayerVisionBound(pos commonmodel.Position) commonmodel.Bound {
	fromX := pos.GetX() - 50
	toX := pos.GetX() + 50

	fromY := pos.GetZ() - 50
	toY := pos.GetZ() + 50

	from := commonmodel.NewPosition(fromX, fromY)
	to := commonmodel.NewPosition(toX, toY)
	bound, _ := commonmodel.NewBound(from, to)

	return bound
}

type Player struct {
	id           commonmodel.PlayerId      // Id of the player
	worldId      commonmodel.WorldId       // The id of the world the player belongs to
	name         string                    // The name of the player
	position     commonmodel.Position      // The current position of the player
	direction    commonmodel.Direction     // The direction where the player is facing
	heldItemId   *commonmodel.ItemId       // Optional, The item held by the player
	domainEvents []domainmodel.DomainEvent // Domain events
	createdAt    time.Time
	updatedAt    time.Time
}

// Interface Implementation Check
var _ domainmodel.Aggregate = (*Player)(nil)

func NewPlayer(
	id commonmodel.PlayerId,
	worldId commonmodel.WorldId,
	name string,
	position commonmodel.Position,
	direction commonmodel.Direction,
	heldItemId *commonmodel.ItemId,
) Player {
	player := Player{
		id:           id,
		worldId:      worldId,
		name:         name,
		position:     position,
		direction:    direction,
		heldItemId:   heldItemId,
		domainEvents: []domainmodel.DomainEvent{},
		createdAt:    time.Now(),
		updatedAt:    time.Now(),
	}
	return player
}

func LoadPlayer(
	id commonmodel.PlayerId,
	worldId commonmodel.WorldId,
	name string,
	position commonmodel.Position,
	direction commonmodel.Direction,
	heldItemId *commonmodel.ItemId,
	createdAt time.Time,
	updatedAt time.Time,
) Player {
	player := Player{
		id:           id,
		worldId:      worldId,
		name:         name,
		position:     position,
		direction:    direction,
		heldItemId:   heldItemId,
		domainEvents: []domainmodel.DomainEvent{},
		createdAt:    createdAt,
		updatedAt:    updatedAt,
	}
	return player
}

func (player *Player) AddDomainEvent(domainEvent domainmodel.DomainEvent) {
	player.domainEvents = append(player.domainEvents, domainEvent)
}

func (player *Player) GetDomainEvents() []domainmodel.DomainEvent {
	return player.domainEvents
}

func (player *Player) GetId() commonmodel.PlayerId {
	return player.id
}

func (player *Player) GetWorldId() commonmodel.WorldId {
	return player.worldId
}

func (player *Player) GetName() string {
	return player.name
}

func (player *Player) GetPosition() commonmodel.Position {
	return player.position
}

func (player *Player) ChangePosition(position commonmodel.Position) {
	player.position = position
}

func (player *Player) GetDirection() commonmodel.Direction {
	return player.direction
}

func (player *Player) ChangeDirection(direction commonmodel.Direction) {
	player.direction = direction
}

func (player *Player) GetVisionBound() commonmodel.Bound {
	return calculatePlayerVisionBound(player.position)
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
