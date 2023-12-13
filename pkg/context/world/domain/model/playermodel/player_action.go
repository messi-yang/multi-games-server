package playermodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

type PlayerAction struct {
	name      PlayerActionNameEnum
	position  worldcommonmodel.Position
	direction worldcommonmodel.Direction
	time      time.Time
}

// Interface Implementation Check
var _ domain.ValueObject[PlayerAction] = (*PlayerAction)(nil)

func NewPlayerAction(name PlayerActionNameEnum, position worldcommonmodel.Position, direction worldcommonmodel.Direction, time time.Time) PlayerAction {
	return PlayerAction{
		name:      name,
		position:  position,
		direction: direction,
		time:      time,
	}
}

func (playerAction PlayerAction) IsEqual(otherPlayerAction PlayerAction) bool {
	return playerAction.name == otherPlayerAction.name
}

func (playerAction PlayerAction) String() string {
	return string(playerAction.name)
}

func (playerAction PlayerAction) GetName() PlayerActionNameEnum {
	return playerAction.name
}

func (playerAction PlayerAction) GetPosition() worldcommonmodel.Position {
	return playerAction.position
}

func (playerAction PlayerAction) UpdatePosition(position worldcommonmodel.Position) PlayerAction {
	return NewPlayerAction(playerAction.name, position, playerAction.direction, playerAction.time)
}

func (playerAction PlayerAction) GetDirection() worldcommonmodel.Direction {
	return playerAction.direction
}

func (playerAction PlayerAction) UpdateDirection(direction worldcommonmodel.Direction) PlayerAction {
	return NewPlayerAction(playerAction.name, playerAction.position, direction, playerAction.time)
}

func (playerAction PlayerAction) GetTime() time.Time {
	return playerAction.time
}

func (playerAction PlayerAction) UpdateTime(time time.Time) PlayerAction {
	return NewPlayerAction(playerAction.name, playerAction.position, playerAction.direction, time)
}
