package playermodel

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
)

type playerActionValue string

const (
	playerActionStand playerActionValue = "stand"
	playerActionWalk  playerActionValue = "walk"
)

type PlayerAction struct {
	action playerActionValue
}

// Interface Implementation Check
var _ domain.ValueObject[PlayerAction] = (*PlayerAction)(nil)

func NewPlayerAction(value string) (PlayerAction, error) {
	switch value {
	case "stand":
		return PlayerAction{
			action: playerActionStand,
		}, nil
	case "walk":
		return PlayerAction{
			action: playerActionWalk,
		}, nil
	default:
		return PlayerAction{}, fmt.Errorf("invalid player action: %s", value)
	}
}

func NewPlayerActionStand() PlayerAction {
	return PlayerAction{
		action: playerActionStand,
	}
}

func NewPlayerActionWalk() PlayerAction {
	return PlayerAction{
		action: playerActionWalk,
	}
}

func (playerAction PlayerAction) IsEqual(otherPlayerAction PlayerAction) bool {
	return playerAction.action == otherPlayerAction.action
}

func (playerAction PlayerAction) String() string {
	return string(playerAction.action)
}
