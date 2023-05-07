package seedclihandler

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/dbseedappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/messaging/memdomainevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/pguow"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (handler *Handler) Exec() {
	pgUow := pguow.NewUow()

	itemRepo := pgrepo.NewItemRepo(pgUow)
	domainEventDispatcher := memdomainevent.NewDispatcher(pgUow)
	dbSeedAppService := dbseedappsrv.NewService(itemRepo, domainEventDispatcher)

	fmt.Println("Start seeding Postgres database")
	err := dbSeedAppService.AddDefaultItems()
	if err != nil {
		pgUow.RevertChanges()
		panic(err)
	}
	pgUow.SaveChanges()
	fmt.Println("Finished seeding Postgres database")
}
