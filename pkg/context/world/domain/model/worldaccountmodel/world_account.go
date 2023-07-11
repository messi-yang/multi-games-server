package worldaccountmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/google/uuid"
)

type WorldAccount struct {
	id                   WorldAccountId
	userId               sharedkernelmodel.UserId
	worldsCount          int8
	worldsCountLimit     int8
	domainEventCollector *domain.DomainEventCollector
}

// Interface Implementation Check
var _ domain.Aggregate = (*WorldAccount)(nil)

func NewWorldAccount(
	userId sharedkernelmodel.UserId,
) WorldAccount {
	return WorldAccount{
		id:                   NewWorldAccountId(uuid.New()),
		userId:               userId,
		worldsCount:          0,
		worldsCountLimit:     1,
		domainEventCollector: domain.NewDomainEventCollector(),
	}
}

func LoadPlayer(
	id WorldAccountId,
	userId sharedkernelmodel.UserId,
	worldsCount int8,
	worldsCountLimit int8,
) WorldAccount {
	return WorldAccount{
		id:                   id,
		userId:               userId,
		worldsCount:          worldsCount,
		worldsCountLimit:     worldsCountLimit,
		domainEventCollector: domain.NewDomainEventCollector(),
	}
}

func (worldAccount *WorldAccount) PopDomainEvents() []domain.DomainEvent {
	return worldAccount.domainEventCollector.PopAll()
}

func (worldAccount *WorldAccount) GetId() WorldAccountId {
	return worldAccount.id
}

func (worldAccount *WorldAccount) GetUserId() sharedkernelmodel.UserId {
	return worldAccount.userId
}

func (worldAccount *WorldAccount) GetWorldsCount() int8 {
	return worldAccount.worldsCount
}

func (worldAccount *WorldAccount) AddWorldsCount() {
	worldAccount.worldsCount += 1
}

func (worldAccount *WorldAccount) GetWorldsCountLimit() int8 {
	return worldAccount.worldsCountLimit
}

func (worldAccount *WorldAccount) CanAddNewWorld() bool {
	return worldAccount.GetWorldsCount() < worldAccount.GetWorldsCountLimit()
}
