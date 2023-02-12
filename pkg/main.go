package main

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/service/gamesocketappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/infrastructure/messaging/redispub"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/infrastructure/persistence/memrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/api/gamesocketapi"

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

	gameSocketAppService := gamesocketappservice.NewService(intEventPublisher, gameRepo, unitRepo, itemRepo)
	gameSocketApiController := gamesocketapi.NewController(gameSocketAppService)

	gameSocketAppService.CreateGame("20716447-6514-4eac-bd05-e558ca72bf3c")

	router.Static("/assets", "./pkg/interface/assets")
	router.Group("/ws/game").GET("/", gameSocketApiController.HandleGameConnection)
	router.Run()

}
