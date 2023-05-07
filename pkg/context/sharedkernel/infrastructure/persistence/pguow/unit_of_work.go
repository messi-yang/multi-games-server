package pguow

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/pgclient"
	"gorm.io/gorm"
)

type Uow interface {
	Execute(func(transaction *gorm.DB) error) error
	RevertChanges()
	SaveChanges()
}

type uow struct {
	transaction *gorm.DB
}

// Dummy Unit of Work, by using this, you don't have to
// to RevertChanges and SaveChanges in the end because it uses a fake transaction.
func NewDummyUow() Uow {
	pgClient := pgclient.GetPgClient()

	return &uow{
		transaction: pgClient,
	}
}

func NewUow() Uow {
	pgClient := pgclient.GetPgClient()

	transaction := pgClient.Begin()
	return &uow{
		transaction: transaction,
	}
}

func (uow *uow) Execute(execute func(transaction *gorm.DB) error) error {
	return execute(uow.transaction)
}

func (uow *uow) RevertChanges() {
	uow.transaction.Rollback()
}

func (uow *uow) SaveChanges() {
	uow.transaction.Commit()
}
