package worldapi

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/service/worldapiservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/transport/api"
	"github.com/gin-gonic/gin"
)

func queryWorldHandler(c *gin.Context) {
	httpPresenter := api.NewHttpPresenter(c)
	worldAppService, err := newWorldAppService(httpPresenter)
	if err != nil {
		httpPresenter.OnError(err)
		return
	}

	worldAppService.QueryWorlds()
}

func createWorldHandler(c *gin.Context) {
	httpPresenter := api.NewHttpPresenter(c)

	var requestDto worldapiservice.CreateWorldRequestDto
	if err := c.BindJSON(&requestDto); err != nil {
		httpPresenter.OnError(err)
		return
	}

	worldAppService, err := newWorldAppService(httpPresenter)
	if err != nil {
		httpPresenter.OnError(err)
		return
	}

	worldAppService.CreateWorld(requestDto.UserId)
}
