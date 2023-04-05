package worldhttpcontroller

import (
	"net/http"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/service/worldappservice"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getWorldHandler(c *gin.Context) {
	worldIdDto, err := uuid.Parse(c.Param("worldId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	worldAppService, err := provideWorldAppService()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	worldDto, err := worldAppService.GetWorld(worldappservice.GetWorldQuery{WorldId: worldIdDto})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, getWorldResponseDto(worldDto))
}

func getWorldsHandler(c *gin.Context) {
	worldAppService, err := provideWorldAppService()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	worldDtos, err := worldAppService.GetWorlds(worldappservice.GetWorldsQuery{})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, getWorldsResponseDto(worldDtos))
}

func createWorldHandler(c *gin.Context) {
	worldAppService, err := provideWorldAppService()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var requestDto createWorldRequestDto
	if err := c.BindJSON(&requestDto); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newWorldIdDto, err := worldAppService.CreateWorld(worldappservice.CreateWorldCommand{UserId: requestDto.UserId})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	worldDto, err := worldAppService.GetWorld(worldappservice.GetWorldQuery{WorldId: newWorldIdDto})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, createWorldResponseDto(worldDto))
}
