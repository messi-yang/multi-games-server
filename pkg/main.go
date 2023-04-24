package main

import (
	"flag"
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/dbseedappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/cli/seedclihandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/assethttphandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/authhttphandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/gamerhttphandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/gamesockethandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/itemhttphandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/worldhttphandler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		switch args[0] {
		case "db-seed":
			itemRepository, err := pgrepo.NewItemRepository()
			if err != nil {
				panic(err)
			}
			dbSeedAppService := dbseedappsrv.NewService(itemRepository)
			seedCliHandler := seedclihandler.NewHandler(dbSeedAppService)
			seedCliHandler.Exec()
			os.Exit(0)
		}
		return
	}

	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))

	assethttphandler.Setup(router)
	gamesockethandler.Setup(router)
	worldhttphandler.Setup(router)
	itemhttphandler.Setup(router)
	gamerhttphandler.Setup(router)
	authhttphandler.Setup(router)
	err := router.Run()
	if err != nil {
		panic(err)
	}
}
