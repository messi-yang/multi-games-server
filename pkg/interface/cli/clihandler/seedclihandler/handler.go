package seedclihandler

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/dbseedappsrv"
)

type Handler struct {
	dbSeedAppService dbseedappsrv.Service
}

var handlerSingleton *Handler

func NewHandler(
	dbSeedAppService dbseedappsrv.Service,
) *Handler {
	if handlerSingleton != nil {
		return handlerSingleton
	}
	return &Handler{dbSeedAppService: dbSeedAppService}
}

func (handler *Handler) Exec() {
	fmt.Println("Start seeding Postgres database")
	err := handler.dbSeedAppService.AddDefaultItems()
	if err != nil {
		panic(err)
	}
	fmt.Println("Finished seeding Postgres database")
}
