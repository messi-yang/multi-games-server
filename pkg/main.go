package main

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/appservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/newappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/viewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/infrastructure/memrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/infrastructure/redispub"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/httpcontroller"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/inteventcontroller"
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
	newGameAppService := newappservice.NewGameAppService(
		gameRepo,
		unitRepo,
		itemRepo,
		intEventPublisher,
	)
	itemAppService := appservice.NewItemAppService(itemRepo)

	itemController := httpcontroller.NewItemHttpController(itemAppService)
	gameController := socketcontroller.NewGameSocketController(
		gameAppService,
	)

	mapSize, _ := commonmodel.NewSizeVo(200, 200)
	newGameAppService.LoadGame(viewmodel.NewSizeVm(mapSize), "20716447-6514-4eac-bd05-e558ca72bf3c")

	go inteventcontroller.NewGameIntEventController(newGameAppService)

	router.Static("/assets", "./pkg/assets")
	router.Group("/ws/game").GET("/", gameController.HandleGameConnection)
	router.GET("/items", itemController.GetAllHandler)
	router.Run()

}
