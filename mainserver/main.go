package mainserver

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/mainserver/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/mainserver/interface/http/gamehandler"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	liveGameApplicationService, _ := applicationservice.NewLiveGameApplicationService(applicationservice.WithLiveGameService())

	router.Group("/ws/game").GET("/", gamehandler.NewHandler(gamehandler.HandlerConfiguration{
		LiveGameApplicationService: liveGameApplicationService,
	}))

	router.Run()
}
