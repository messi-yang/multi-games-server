package seedclihandler

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/event/memory/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/dbseedappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/postgres/pgrepo"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (handler *Handler) Exec() {
	pgUow := pguow.NewUow()

	domainEventDispatcher := memdomainevent.NewDispatcher(pgUow)
	itemRepo := pgrepo.NewItemRepo(pgUow, domainEventDispatcher)
	dbSeedAppService := dbseedappsrv.NewService(itemRepo)

	fmt.Println("Start seeding Postgres database")
	err := dbSeedAppService.AddDefaultItems()
	if err != nil {
		pgUow.RevertChanges()
		panic(err)
	}
	pgUow.SaveChanges()
	fmt.Println("Finished seeding Postgres database")
}
