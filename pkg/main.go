package main

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/appservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/infrastructure/memrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/infrastructure/redispub"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/socketcontroller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))

	intEventPublisher := redispub.New()

	itemRepo := memrepo.NewItemMemRepo()
	gameRepo := memrepo.NewGameMemRepo()
	unitRepo := memrepo.NewUnitMemRepo()

	gameAppService := appservice.NewGameAppService(intEventPublisher, gameRepo, unitRepo, itemRepo)
	gameController := socketcontroller.NewGameSocketController(
		gameAppService,
	)

	gameAppService.LoadGame("20716447-6514-4eac-bd05-e558ca72bf3c")

	router.Static("/assets", "./pkg/assets")
	router.Group("/ws/game").GET("/", gameController.HandleGameConnection)
	router.Run()

}
