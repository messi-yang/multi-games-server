package main

import (
	"flag"
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/cli/seedclicontroller"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/assethttpcontroller"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/gamerhttpcontroller"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/gamesocketcontroller"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/itemhttpcontroller"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/worldhttpcontroller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		switch args[0] {
		case "db-seed":
			seedclicontroller.Exec()
			os.Exit(0)
		}
		return
	}

	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))

	assethttpcontroller.Setup(router)
	gamesocketcontroller.Setup(router)
	worldhttpcontroller.Setup(router)
	itemhttpcontroller.Setup(router)
	gamerhttpcontroller.Setup(router)
	err := router.Run()
	if err != nil {
		panic(err)
	}
}
