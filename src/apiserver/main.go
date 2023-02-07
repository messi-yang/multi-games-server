package main

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/application/appservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/interface/httpcontroller"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/interface/socketcontroller"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/infrastructure/commonmemrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/infrastructure/redispub"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))

	intEventPublisher := redispub.New()
	itemRepo := commonmemrepo.NewItemMemRepo()
	gameAppService := appservice.NewGameAppService(intEventPublisher, itemRepo)
	itemAppService := appservice.NewItemAppService(itemRepo)

	itemController := httpcontroller.NewItemHttpController(itemAppService)
	gameController := socketcontroller.NewGameSocketController(
		gameAppService,
	)

	router.Static("/assets", "./src/assets")

	router.Group("/ws/game").GET("/", gameController.HandleGameConnection)
	router.GET("/items", itemController.GetAllHandler)

	router.Run()
}
