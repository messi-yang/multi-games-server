package seedclihandler

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/dbseedappservice"
)

type Handler struct {
	dbSeedAppService dbseedappservice.Service
}

var handlerSingleton *Handler

func NewHandler(
	dbSeedAppService dbseedappservice.Service,
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
