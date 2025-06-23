package roommodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/gamemodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/google/uuid"
)

type Room struct {
	id                   globalcommonmodel.RoomId
	userId               globalcommonmodel.UserId
	name                 string
	currentGameId        *gamemodel.GameId
	createdAt            time.Time
	updatedAt            time.Time
	domainEventCollector *domain.DomainEventCollector
}

// Interface Implementation Check
var _ domain.DomainEventDispatchableAggregate = (*Room)(nil)

func NewRoom(
	userId globalcommonmodel.UserId,
	currentGameId *gamemodel.GameId,
	name string,
) Room {
	newRoom := Room{
		id:                   globalcommonmodel.NewRoomId(uuid.New()),
		userId:               userId,
		name:                 name,
		currentGameId:        currentGameId,
		createdAt:            time.Now(),
		updatedAt:            time.Now(),
		domainEventCollector: domain.NewDomainEventCollector(),
	}
	newRoom.domainEventCollector.Add(NewRoomCreated(
		newRoom.id,
		newRoom.userId,
	))
	return newRoom
}

func LoadRoom(
	roomId globalcommonmodel.RoomId,
	userId globalcommonmodel.UserId,
	name string,
	currentGameId *gamemodel.GameId,
	createdAt time.Time,
	updatedAt time.Time,
) Room {
	return Room{
		id:                   roomId,
		userId:               userId,
		name:                 name,
		currentGameId:        currentGameId,
		createdAt:            createdAt,
		updatedAt:            updatedAt,
		domainEventCollector: domain.NewDomainEventCollector(),
	}
}

func (room *Room) PopDomainEvents() []domain.DomainEvent {
	return room.domainEventCollector.PopAll()
}

func (room *Room) GetId() globalcommonmodel.RoomId {
	return room.id
}

func (room *Room) GetUserId() globalcommonmodel.UserId {
	return room.userId
}

func (room *Room) GetName() string {
	return room.name
}

func (room *Room) ChangeName(name string) {
	room.name = name
}

func (room *Room) GetCurrentGameId() *gamemodel.GameId {
	return room.currentGameId
}

func (room *Room) SetCurrentGameId(currentGameId *gamemodel.GameId) {
	room.currentGameId = currentGameId
}

func (room *Room) GetCreatedAt() time.Time {
	return room.createdAt
}

func (room *Room) GetUpdatedAt() time.Time {
	return room.updatedAt
}

func (room *Room) Delete() {
	room.domainEventCollector.Add(NewRoomDeleted(
		room.id,
		room.userId,
	))
}
