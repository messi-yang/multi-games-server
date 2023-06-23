package worldaccountmodel

import (
	"errors"

	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/google/uuid"
)

var (
	ErrWorldsCountExceedsLimit = errors.New("worlds count has reached the limit")
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
	worldsCount int8,
	worldsCountLimit int8,
) WorldAccount {
	return WorldAccount{
		id:                   NewWorldAccountId(uuid.New()),
		userId:               userId,
		worldsCount:          worldsCount,
		worldsCountLimit:     worldsCountLimit,
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

func (worldAccount *WorldAccount) GetWorldsCountLimit() int8 {
	return worldAccount.worldsCountLimit
}

func (worldAccount *WorldAccount) AddWorldsCount() error {
	if worldAccount.GetWorldsCount() >= worldAccount.GetWorldsCountLimit() {
		return ErrWorldsCountExceedsLimit
	}
	worldAccount.worldsCount += 1
	return nil
}
