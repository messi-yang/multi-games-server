package worldhttphandler

import (
	"net/http"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/worldappservice"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type httpHandler struct {
	worldAppService worldappservice.Service
}

var httpHandlerSingleton *httpHandler

func newHttpHandler(
	worldAppService worldappservice.Service,
) *httpHandler {
	if httpHandlerSingleton != nil {
		return httpHandlerSingleton
	}
	return &httpHandler{worldAppService: worldAppService}
}

func (httpHandler *httpHandler) getWorld(c *gin.Context) {
	worldIdDto, err := uuid.Parse(c.Param("worldId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	worldDto, err := httpHandler.worldAppService.GetWorld(worldappservice.GetWorldQuery{WorldId: worldIdDto})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, getWorldResponseDto(worldDto))
}

func (httpHandler *httpHandler) queryWorlds(c *gin.Context) {
	worldDtos, err := httpHandler.worldAppService.QueryWorlds(worldappservice.QueryWorldsQuery{})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, queryWorldsResponseDto(worldDtos))
}

func (httpHandler *httpHandler) createWorld(c *gin.Context) {

	var requestDto createWorldRequestDto
	if err := c.BindJSON(&requestDto); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newWorldIdDto, err := httpHandler.worldAppService.CreateWorld(worldappservice.CreateWorldCommand{GamerId: requestDto.GamerId})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	worldDto, err := httpHandler.worldAppService.GetWorld(worldappservice.GetWorldQuery{WorldId: newWorldIdDto})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, createWorldResponseDto(worldDto))
}
