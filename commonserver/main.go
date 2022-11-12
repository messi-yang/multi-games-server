package commonserver

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/commonserver/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/commonserver/interface/http/gamehandler"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	gameApplicationService, _ := applicationservice.NewGameApplicationService(applicationservice.WithGameService())

	router.Group("/ws/game").GET("/", gamehandler.NewHandler(gamehandler.HandlerConfiguration{
		GameApplicationService: gameApplicationService,
	}))

	router.Run()
}
