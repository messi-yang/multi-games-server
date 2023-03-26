package worldapi

import (
	"net/http"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/dto"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

func queryWorldHandler(c *gin.Context) {
	worldAppService, err := newWorldAppService()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	worldAppService.QueryWorlds(
		func(worlds []worldmodel.WorldAgg) {
			worldDtos := lo.Map(worlds, func(world worldmodel.WorldAgg, _ int) dto.WorldAggDto {
				return dto.NewWorldAggDto(world)
			})
			c.JSON(http.StatusOK, queryWorldsResponseDto(worldDtos))
		},
		func(err error) {
			c.JSON(http.StatusBadRequest, err.Error())
		},
	)
}

func createWorldHandler(c *gin.Context) {
	worldAppService, err := newWorldAppService()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var requestDto createWorldRequestDto
	if err := c.BindJSON(&requestDto); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	worldAppService.CreateWorld(
		requestDto.UserId,
		func(world worldmodel.WorldAgg) {
			worldDto := dto.NewWorldAggDto(world)
			c.JSON(http.StatusOK, createWorldResponseDto(worldDto))
		},
		func(err error) {
			c.JSON(http.StatusBadRequest, err.Error())
		},
	)
}
