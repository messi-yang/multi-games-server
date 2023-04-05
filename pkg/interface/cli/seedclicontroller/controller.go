package seedclicontroller

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/dbseedappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/postgres"
)

func Exec() {
	fmt.Println("Start seeding Postgres database")

	itemRepository, err := postgres.NewItemRepository()
	if err != nil {
		panic(err)
	}

	dbSeedAppService := dbseedappservice.NewService(itemRepository)
	err = dbSeedAppService.AddDefaultItems()
	if err != nil {
		panic(err)
	}

	fmt.Println("Finished seeding Postgres database")
}
