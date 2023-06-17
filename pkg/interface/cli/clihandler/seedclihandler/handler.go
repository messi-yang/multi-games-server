package seedclihandler

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/game/application/service/dbseedappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/infrastructure/persistence/postgres/pgrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/event/memory/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
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
