package seedcmd

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/service/dbseedcmdservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/infrastructure/persistence/postgres"
)

func Exec() {
	fmt.Println("Start seeding Postgres database")

	itemRepository, err := postgres.NewItemRepository()
	if err != nil {
		panic(err)
	}

	dbSeedAppService := dbseedcmdservice.NewService(itemRepository)
	err = dbSeedAppService.AddDefaultItems()
	if err != nil {
		panic(err)
	}

	fmt.Println("Finished seeding Postgres database")
}
