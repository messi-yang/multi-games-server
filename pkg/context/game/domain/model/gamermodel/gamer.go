package gamermodel

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/google/uuid"
)

var (
	ErrWorldsCountExceedsLimit = errors.New("worlds count has reached the limit")
)

type Gamer struct {
	id                   GamerId
	userId               sharedkernelmodel.UserId
	worldsCount          int8
	worldsCountLimit     int8
	domainEventCollector *domain.DomainEventCollector
}

// Interface Implementation Check
var _ domain.Aggregate = (*Gamer)(nil)

func NewGamer(
	userId sharedkernelmodel.UserId,
	worldsCount int8,
	worldsCountLimit int8,
) Gamer {
	return Gamer{
		id:                   NewGamerId(uuid.New()),
		userId:               userId,
		worldsCount:          worldsCount,
		worldsCountLimit:     worldsCountLimit,
		domainEventCollector: domain.NewDomainEventCollector(),
	}
}

func LoadPlayer(
	id GamerId,
	userId sharedkernelmodel.UserId,
	worldsCount int8,
	worldsCountLimit int8,
) Gamer {
	return Gamer{
		id:                   id,
		userId:               userId,
		worldsCount:          worldsCount,
		worldsCountLimit:     worldsCountLimit,
		domainEventCollector: domain.NewDomainEventCollector(),
	}
}

func (gamer *Gamer) PopDomainEvents() []domain.DomainEvent {
	return gamer.domainEventCollector.PopAll()
}

func (gamer *Gamer) GetId() GamerId {
	return gamer.id
}

func (gamer *Gamer) GetUserId() sharedkernelmodel.UserId {
	return gamer.userId
}

func (gamer *Gamer) GetWorldsCount() int8 {
	return gamer.worldsCount
}

func (gamer *Gamer) GetWorldsCountLimit() int8 {
	return gamer.worldsCountLimit
}

func (gamer *Gamer) AddWorldsCount() error {
	if gamer.GetWorldsCount() >= gamer.GetWorldsCountLimit() {
		return ErrWorldsCountExceedsLimit
	}
	gamer.worldsCount += 1
	return nil
}
