package worldhttphandler

import (
	"net/http"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/gamerappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/worldappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/application/service/identityappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/httputil"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HttpHandler struct {
	identityAppService identityappsrv.Service
	worldAppService    worldappsrv.Service
	gamerAppService    gamerappsrv.Service
}

var httpHandlerSingleton *HttpHandler

func NewHttpHandler(
	identityAppService identityappsrv.Service,
	worldAppService worldappsrv.Service,
	gamerAppService gamerappsrv.Service,
) *HttpHandler {
	if httpHandlerSingleton != nil {
		return httpHandlerSingleton
	}
	return &HttpHandler{
		identityAppService: identityAppService,
		worldAppService:    worldAppService,
		gamerAppService:    gamerAppService,
	}
}

func (httpHandler *HttpHandler) GetWorld(c *gin.Context) {
	worldIdDto, err := uuid.Parse(c.Param("worldId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	worldDto, err := httpHandler.worldAppService.GetWorld(worldappsrv.GetWorldQuery{WorldId: worldIdDto})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, getWorldResponseDto(worldDto))
}

func (httpHandler *HttpHandler) QueryWorlds(c *gin.Context) {
	worldDtos, err := httpHandler.worldAppService.QueryWorlds(worldappsrv.QueryWorldsQuery{})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, queryWorldsResponseDto(worldDtos))
}

func (httpHandler *HttpHandler) CreateWorld(c *gin.Context) {
	userIdDto := httputil.GetUserId(c)
	gamer, err := httpHandler.gamerAppService.GetGamerByUserId(gamerappsrv.GetGamerByUserIdQuery{UserId: userIdDto})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newWorldIdDto, err := httpHandler.worldAppService.CreateWorld(worldappsrv.CreateWorldCommand{GamerId: gamer.Id})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	worldDto, err := httpHandler.worldAppService.GetWorld(worldappsrv.GetWorldQuery{WorldId: newWorldIdDto})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, createWorldResponseDto(worldDto))
}
