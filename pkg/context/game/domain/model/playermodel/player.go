package playermodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/google/uuid"
)

type Player struct {
	id           PlayerId                 // Id of the player
	roomId       globalcommonmodel.RoomId // The id of the game the player belongs to
	userId       *globalcommonmodel.UserId
	name         string // The name of the playe
	hostPriority float64
	createdAt    time.Time
	updatedAt    time.Time
}

// Interface Implementation Check
var _ domain.Aggregate = (*Player)(nil)

func NewPlayer(
	roomId globalcommonmodel.RoomId,
	userId *globalcommonmodel.UserId,
	name string,
	hostPriority float64,
) Player {
	return Player{
		id:           NewPlayerId(uuid.New()),
		roomId:       roomId,
		userId:       userId,
		name:         name,
		hostPriority: hostPriority,
		createdAt:    time.Now(),
		updatedAt:    time.Now(),
	}
}

func LoadPlayer(
	id PlayerId,
	roomId globalcommonmodel.RoomId,
	userId *globalcommonmodel.UserId,
	name string,
	hostPriority float64,
	createdAt time.Time,
	updatedAt time.Time,
) Player {
	player := Player{
		id:           id,
		roomId:       roomId,
		userId:       userId,
		name:         name,
		hostPriority: hostPriority,
		createdAt:    createdAt,
		updatedAt:    updatedAt,
	}
	return player
}

func (player *Player) GetId() PlayerId {
	return player.id
}

func (player *Player) GetRoomId() globalcommonmodel.RoomId {
	return player.roomId
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

func (player *Player) GetHostPriority() float64 {
	return player.hostPriority
}

func (player *Player) SetHostPriority(hostPriority float64) {
	player.hostPriority = hostPriority
}

func (player *Player) GetCreatedAt() time.Time {
	return player.createdAt
}

func (player *Player) GetUpdatedAt() time.Time {
	return player.updatedAt
}
