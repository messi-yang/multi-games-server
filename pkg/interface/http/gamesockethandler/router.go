package gamesockethandler

import (
	"net/http"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/gameappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/memrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/pgrepository"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func Setup(router *gin.Engine) {
	itemRepository, err := pgrepository.NewItemRepository()
	if err != nil {
		panic(err)
	}
	playerRepository := memrepo.NewPlayerMemRepository()
	worldRepository, err := pgrepository.NewWorldRepository()
	if err != nil {
		panic(err)
	}
	unitRepository, err := pgrepository.NewUnitRepository()
	if err != nil {
		panic(err)
	}
	gameService := service.NewGameService(worldRepository, playerRepository, unitRepository, itemRepository)
	gameAppService := gameappservice.NewService(
		worldRepository, playerRepository, unitRepository, itemRepository, gameService,
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
