package pguow

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/application/uow"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/pgclient"
	"gorm.io/gorm"
)

type Uow struct {
	transaction *gorm.DB
}

// Interface Implementation Check
var _ uow.Uow[*gorm.DB] = (*Uow)(nil)

// Dummy Unit of Work, by using this, you don't have to
// to Rollback and Commit in the end because it uses a fake transaction.
func NewDummyUow() *Uow {
	pgClient := pgclient.NewPgClient()

	return &Uow{
		transaction: pgClient,
	}
}

func NewUow() *Uow {
	pgClient := pgclient.NewPgClient()

	transaction := pgClient.Begin()
	return &Uow{
		transaction: transaction,
	}
}

func (uow *Uow) GetTransaction() *gorm.DB {
	return uow.transaction
}

func (uow *Uow) Rollback() {
	uow.transaction.Rollback()
}

func (uow *Uow) Commit() {
	uow.transaction.Commit()
}
