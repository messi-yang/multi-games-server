package main

import (
	"flag"
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/dbseedappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/cli/seedclihandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/httprouter"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
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
		return
	}

	err := httprouter.Run()
	if err != nil {
		panic(err)
	}
}
