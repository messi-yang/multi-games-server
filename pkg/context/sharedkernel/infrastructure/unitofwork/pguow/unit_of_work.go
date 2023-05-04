package pguow

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/pgclient"
	"gorm.io/gorm"
)

type Uow interface {
	GetTransaction() *gorm.DB
	DispatchDomainEvents(aggregate domain.Aggregate) error
	Rollback()
	Commit()
}

type uow struct {
	transaction           *gorm.DB
	domainEventDispatcher DomainEventDispatcher
}

// Dummy Unit of Work, by using this, you don't have to
// to Rollback and Commit in the end because it uses a fake transaction.
func NewDummyUow() Uow {
	pgClient := pgclient.GetPgClient()

	return &uow{
		transaction:           pgClient,
		domainEventDispatcher: nil,
	}
}

func NewUow() Uow {
	pgClient := pgclient.GetPgClient()
	domainEventDispatcher := GetDomainEventDispatcher()

	transaction := pgClient.Begin()
	return &uow{
		transaction:           transaction,
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (uow *uow) GetTransaction() *gorm.DB {
	return uow.transaction
}

func (uow *uow) DispatchDomainEvents(aggregate domain.Aggregate) error {
	for _, domainEvent := range aggregate.PopDomainEvents() {
		err := uow.domainEventDispatcher.Dispatch(uow, domainEvent)
		if err != nil {
			return err
		}
	}
	return nil
}

func (uow *uow) Rollback() {
	uow.transaction.Rollback()
}

func (uow *uow) Commit() {
	uow.transaction.Commit()
}
