package seedclihandler

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/dbseedappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/unitofwork/pguow"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (handler *Handler) Exec() {
	pgUow := pguow.NewUow()

	itemRepo := pgrepo.NewItemRepo(pgUow)
	dbSeedAppService := dbseedappsrv.NewService(itemRepo)

	fmt.Println("Start seeding Postgres database")
	err := dbSeedAppService.AddDefaultItems()
	if err != nil {
		pgUow.Rollback()
		panic(err)
	}
	pgUow.Commit()
	fmt.Println("Finished seeding Postgres database")
}
