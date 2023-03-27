package worldapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func queryWorldHandler(c *gin.Context) {
	worldAppService, err := provideWorldAppService()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	worldDtos, err := worldAppService.QueryWorlds()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, queryWorldsResponseDto(worldDtos))
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

	worldDto, err := worldAppService.CreateWorld(requestDto.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, createWorldResponseDto(worldDto))
}
