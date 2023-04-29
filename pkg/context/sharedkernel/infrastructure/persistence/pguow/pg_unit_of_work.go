package pguow

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/pgclient"
	"gorm.io/gorm"
)

type Uow interface {
	GetTransaction() *gorm.DB
	Rollback()
	Commit()
}

type uow struct {
	transaction *gorm.DB
}

// Dummy Unit of Work, by using this, you don't have to
// to Rollback and Commit in the end because it uses a fake transaction.
func NewDummyUow() Uow {
	pgClient := pgclient.NewPgClient()

	return &uow{
		transaction: pgClient,
	}
}

func NewUow() Uow {
	pgClient := pgclient.NewPgClient()

	transaction := pgClient.Begin()
	return &uow{
		transaction: transaction,
	}
}

func (uow *uow) GetTransaction() *gorm.DB {
	return uow.transaction
}

func (uow *uow) Rollback() {
	uow.transaction.Rollback()
}

func (uow *uow) Commit() {
	uow.transaction.Commit()
}
