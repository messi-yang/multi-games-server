package apiserver

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

func Start() {
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))

	gameRepo, _ := psqlrepo.NewGamePsqlRepo()
	IntEventPublisher := redispub.New()
	itemRepo := memrepo.NewItemMemRepo()
	liveGameAppService := appservice.NewLiveGameAppService(IntEventPublisher, itemRepo)
	itemAppService := appservice.NewItemAppService(itemRepo)

	playerRepo := memrepo.NewPlayerMemRepo()
	playerAppService := appservice.NewPlayerAppService(playerRepo)

	itemController := httpcontroller.NewItemHttpController(itemAppService)
	liveGameController := socketcontroller.NewController(
		gameRepo,
		liveGameAppService,
		playerRepo,
	)

	playerController := httpcontroller.NewPlayerHttpController(playerAppService)

	router.Static("/assets", "./src/assets")

	router.Group("/ws/game").GET("/", liveGameController.HandleLiveGameConnection)
	router.GET("/items", itemController.GetAllHandler)
	router.GET("/players", playerController.GetAllHandler)

	router.Run()
}
