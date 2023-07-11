package seedclihandler

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/itemappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/providedependency"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (handler *Handler) Exec() {
	fmt.Println("Start seeding Postgres database")

	pgUow := pguow.NewUow()
	itemAppService := providedependency.ProvideItemAppService(pgUow)

	err := itemAppService.CreateDefaultItems(itemappsrv.CreateDefaultItemsCommand{})
	if err != nil {
		pgUow.RevertChanges()
		panic(err)
	}
	pgUow.SaveChanges()
	fmt.Println("Finished seeding Postgres database")
}
