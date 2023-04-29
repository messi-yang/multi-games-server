package clirouter

import (
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/dbseedappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/cli/clihandler/seedclihandler"
)

func Run(args []string) {
	switch args[0] {
	case "db-seed":
		itemRepo, err := pgrepo.NewItemRepo()
		if err != nil {
			panic(err)
		}
		dbSeedAppService := dbseedappsrv.NewService(itemRepo)
		seedCliHandler := seedclihandler.NewHandler(dbSeedAppService)
		seedCliHandler.Exec()
		os.Exit(0)
	}
}
