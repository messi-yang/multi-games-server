package gamerhttphandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/gamerappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/pgrepo"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	gamerRepository, err := pgrepo.NewGamerRepository()
	if err != nil {
		if err != nil {
			panic(err)
		}
	}
	gamerAppService := gamerappsrv.NewService(gamerRepository)
	httpHandler := newHttpHandler(gamerAppService)

	routerGroup := router.Group("/api/gamers")
	routerGroup.GET("/", httpHandler.queryGamers)
}
