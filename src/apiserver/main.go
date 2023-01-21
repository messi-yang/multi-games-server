package main

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/application/appservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/infrastructure/memrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/interface/httpcontroller"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/interface/socketcontroller"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/infrastructure/psqlrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/infrastructure/redispub"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))

	gameRepo, _ := psqlrepo.NewGamePsqlRepo()
	intEventPublisher := redispub.New()
	itemRepo := memrepo.NewItemMemRepo()
	liveGameAppService := appservice.NewLiveGameAppService(intEventPublisher, itemRepo)
	itemAppService := appservice.NewItemAppService(itemRepo)

	gameAppService := appservice.NewGameAppService(gameRepo)

	itemController := httpcontroller.NewItemHttpController(itemAppService)
	liveGameController := socketcontroller.NewLiveGameSocketController(
		gameRepo,
		liveGameAppService,
	)

	gameController := httpcontroller.NewGameHttpController(gameAppService)

	router.Static("/assets", "./src/assets")

	router.Group("/ws/game").GET("/", liveGameController.HandleLiveGameConnection)
	router.GET("/items", itemController.GetAllHandler)
	router.GET("/games", gameController.GetAllHandler)

	router.Run()
}
