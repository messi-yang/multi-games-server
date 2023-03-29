package worldhttpcontroller

import (
	"net/http"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/service/worldappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/usermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/httpdto"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

func queryWorldHandler(c *gin.Context) {
	worldAppService, err := provideWorldAppService()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	worlds, err := worldAppService.FindWorlds(worldappservice.GetWorldsQuery{})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	worldDtos := lo.Map(worlds, func(world worldmodel.WorldAgg, _ int) httpdto.WorldAggDto {
		return httpdto.NewWorldAggDto(world)
	})
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

	userId := usermodel.NewUserIdVo(requestDto.UserId)

	worldId, err := worldAppService.CreateWorld(worldappservice.CreateWorldCommand{UserId: userId})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	world, err := worldAppService.GetWorld(worldappservice.GetWorldQuery{WorldId: worldId})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	worldDto := httpdto.NewWorldAggDto(world)
	c.JSON(http.StatusOK, createWorldResponseDto(worldDto))
}
