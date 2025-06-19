package playermodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/gamecommonmodel"
)

type PlayerAction struct {
	name      PlayerActionNameEnum
	direction gamecommonmodel.Direction
}

// Interface Implementation Check
var _ domain.ValueObject[PlayerAction] = (*PlayerAction)(nil)

func NewPlayerAction(name PlayerActionNameEnum, direction gamecommonmodel.Direction) PlayerAction {
	return PlayerAction{
		name:      name,
		direction: direction,
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

func (playerAction PlayerAction) GetDirection() gamecommonmodel.Direction {
	return playerAction.direction
}

func (playerAction PlayerAction) UpdateDirection(direction gamecommonmodel.Direction) PlayerAction {
	return NewPlayerAction(playerAction.name, direction)
}
