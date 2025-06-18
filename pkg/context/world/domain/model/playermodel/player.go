package playermodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/google/uuid"
)

type Player struct {
	id              PlayerId                  // Id of the player
	worldId         globalcommonmodel.WorldId // The id of the world the player belongs to
	userId          *globalcommonmodel.UserId
	name            string // The name of the player
	action          PlayerAction
	precisePosition worldcommonmodel.PrecisePosition
	createdAt       time.Time
	updatedAt       time.Time
}

// Interface Implementation Check
var _ domain.Aggregate = (*Player)(nil)

func NewPlayer(
	worldId globalcommonmodel.WorldId,
	name string,
	direction worldcommonmodel.Direction,
) Player {
	return Player{
		id:      NewPlayerId(uuid.New()),
		worldId: worldId,
		name:    name,
		action: NewPlayerAction(
			PlayerActionNameEnumStand,
			worldcommonmodel.NewDirection(0),
		),
		precisePosition: worldcommonmodel.NewPrecisePosition(0, 0),
		createdAt:       time.Now(),
		updatedAt:       time.Now(),
	}
}

func LoadPlayer(
	id PlayerId,
	worldId globalcommonmodel.WorldId,
	userId *globalcommonmodel.UserId,
	name string,
	direction worldcommonmodel.Direction,
	action PlayerAction,
	precisePosition worldcommonmodel.PrecisePosition,
	createdAt time.Time,
	updatedAt time.Time,
) Player {
	player := Player{
		id:              id,
		worldId:         worldId,
		userId:          userId,
		name:            name,
		action:          action,
		precisePosition: precisePosition,
		createdAt:       createdAt,
		updatedAt:       updatedAt,
	}
	return player
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

func (player *Player) UpdateName(name string) {
	player.name = name
}

func (player *Player) GetAction() PlayerAction {
	return player.action
}

func (player *Player) GetPrecisePosition() worldcommonmodel.PrecisePosition {
	return player.precisePosition
}

func (player *Player) UpdatePrecisePosition(precisePosition worldcommonmodel.PrecisePosition) {
	player.precisePosition = precisePosition
}

func (player *Player) Teleport(precisePosition worldcommonmodel.PrecisePosition) {
	player.UpdatePrecisePosition(precisePosition)
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
