package worldaccountmodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/google/uuid"
)

type WorldAccount struct {
	id                   WorldAccountId
	userId               globalcommonmodel.UserId
	worldsCount          int8
	worldsCountLimit     int8
	createdAt            time.Time
	updatedAt            time.Time
	domainEventCollector *domain.DomainEventCollector
}

// Interface Implementation Check
var _ domain.Aggregate = (*WorldAccount)(nil)

func NewWorldAccount(
	userId globalcommonmodel.UserId,
) WorldAccount {
	return WorldAccount{
		id:                   NewWorldAccountId(uuid.New()),
		userId:               userId,
		worldsCount:          0,
		worldsCountLimit:     1,
		createdAt:            time.Now(),
		updatedAt:            time.Now(),
		domainEventCollector: domain.NewDomainEventCollector(),
	}
}

func LoadPlayer(
	id WorldAccountId,
	userId globalcommonmodel.UserId,
	worldsCount int8,
	worldsCountLimit int8,
	createdAt time.Time,
	updatedAt time.Time,
) WorldAccount {
	return WorldAccount{
		id:                   id,
		userId:               userId,
		worldsCount:          worldsCount,
		worldsCountLimit:     worldsCountLimit,
		createdAt:            createdAt,
		updatedAt:            updatedAt,
		domainEventCollector: domain.NewDomainEventCollector(),
	}
}

func (worldAccount *WorldAccount) PopDomainEvents() []domain.DomainEvent {
	return worldAccount.domainEventCollector.PopAll()
}

func (worldAccount *WorldAccount) GetId() WorldAccountId {
	return worldAccount.id
}

func (worldAccount *WorldAccount) GetUserId() globalcommonmodel.UserId {
	return worldAccount.userId
}

func (worldAccount *WorldAccount) GetWorldsCount() int8 {
	return worldAccount.worldsCount
}

func (worldAccount *WorldAccount) AddWorldsCount() {
	worldAccount.worldsCount += 1
}

func (worldAccount *WorldAccount) SubtractWorldsCount() {
	worldAccount.worldsCount -= 1
}

func (worldAccount *WorldAccount) GetWorldsCountLimit() int8 {
	return worldAccount.worldsCountLimit
}

func (worldAccount *WorldAccount) CanAddNewWorld() bool {
	return worldAccount.GetWorldsCount() < worldAccount.GetWorldsCountLimit()
}

func (worldAccount *WorldAccount) GetCreatedAt() time.Time {
	return worldAccount.createdAt
}

func (worldAccount *WorldAccount) GetUpdatedAt() time.Time {
	return worldAccount.updatedAt
}
