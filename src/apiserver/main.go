package apiserver

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/application/service/itemappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/application/service/livegameappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/application/service/playerappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/infrastructure/memrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/interface/controller/itemcontroller"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/interface/controller/livegamecontroller"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/interface/controller/playercontroller"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/infrastructure/psqlrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/infrastructure/redispub"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))

	gameRepo, _ := psqlrepo.NewGamePsqlRepo()
	intgrEventPublisher := redispub.New()
	itemRepo := memrepo.NewItemMemRepo()
	liveGameAppService := livegameappservice.New(intgrEventPublisher, itemRepo)
	itemAppService := itemappservice.New(itemRepo)

	playerRepo := memrepo.NewPlayerMemRepo()
	playerAppService := playerappservice.New(playerRepo)

	itemController := itemcontroller.New(itemAppService)
	liveGameController := livegamecontroller.NewController(
		gameRepo,
		liveGameAppService,
		playerRepo,
	)

	playerController := playercontroller.New(playerAppService)

	router.Static("/assets", "./src/assets")

	router.Group("/ws/game").GET("/", liveGameController.HandleLiveGameConnection)
	router.GET("/items", itemController.GetAllHandler)
	router.GET("/players", playerController.GetAllHandler)

	router.Run()
}
