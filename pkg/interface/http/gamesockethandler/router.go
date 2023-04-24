package gamesockethandler

import (
	"net/http"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/gameappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/service/gamedomainsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/memrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/pgrepo"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func Setup(router *gin.Engine) {
	itemRepository, err := pgrepo.NewItemRepository()
	if err != nil {
		panic(err)
	}
	playerRepository := memrepo.NewPlayerMemRepository()
	worldRepository, err := pgrepo.NewWorldRepository()
	if err != nil {
		panic(err)
	}
	unitRepository, err := pgrepo.NewUnitRepository()
	if err != nil {
		panic(err)
	}
	gameDomainService := gamedomainsrv.NewService(worldRepository, playerRepository, unitRepository, itemRepository)
	gameAppService := gameappsrv.NewService(
		worldRepository, playerRepository, unitRepository, itemRepository, gameDomainService,
	)
	httpHandler := newHttpHandler(gameAppService, websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	})

	routerGroup := router.Group("/ws/game")
	routerGroup.GET("/", httpHandler.gameConnection)
}
