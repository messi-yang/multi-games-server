package worldhttphandler

import (
	"net/http"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/gamerappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/worldappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/httputil"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HttpHandler struct{}

func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}

func (httpHandler *HttpHandler) GetWorld(c *gin.Context) {
	worldIdDto, err := uuid.Parse(c.Param("worldId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	pgUow := pguow.NewDummyUow()

	worldAppService := provideWorldAppService(pgUow)
	worldDto, err := worldAppService.GetWorld(worldappsrv.GetWorldQuery{WorldId: worldIdDto})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getWorldResponse(worldDto))
}

func (httpHandler *HttpHandler) QueryWorlds(c *gin.Context) {
	pgUow := pguow.NewDummyUow()

	worldAppService := provideWorldAppService(pgUow)
	worldDtos, err := worldAppService.QueryWorlds(worldappsrv.QueryWorldsQuery{})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, queryWorldsResponse(worldDtos))
}

func (httpHandler *HttpHandler) CreateWorld(c *gin.Context) {
	userIdDto := httputil.GetUserId(c)

	pgUow := pguow.NewUow()

	gamerAppService := provideGamerAppService(pgUow)
	worldAppService := provideWorldAppService(pgUow)

	gamer, err := gamerAppService.GetGamerByUserId(gamerappsrv.GetGamerByUserIdQuery{UserId: userIdDto})
	if err != nil {
		pgUow.RevertChanges()
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newWorldIdDto, err := worldAppService.CreateWorld(worldappsrv.CreateWorldCommand{GamerId: gamer.Id})
	if err != nil {
		pgUow.RevertChanges()
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	worldDto, err := worldAppService.GetWorld(worldappsrv.GetWorldQuery{WorldId: newWorldIdDto})
	if err != nil {
		pgUow.RevertChanges()
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	pgUow.SaveChanges()
	c.JSON(http.StatusOK, createWorldResponse(worldDto))
}

func (httpHandler *HttpHandler) UpdateWorld(c *gin.Context) {
	userIdDto := httputil.GetUserId(c)

	var requestBody updateWorldRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	worldIdDto, err := uuid.Parse(c.Param("worldId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	pgUow := pguow.NewUow()

	gamerAppService := provideGamerAppService(pgUow)
	worldAppService := provideWorldAppService(pgUow)

	gamer, err := gamerAppService.GetGamerByUserId(gamerappsrv.GetGamerByUserIdQuery{UserId: userIdDto})
	if err != nil {
		pgUow.RevertChanges()
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err = worldAppService.UpdateWorld(worldappsrv.UpdateWorldCommand{
		GamerId: gamer.Id,
		WorldId: worldIdDto,
		Name:    requestBody.Name,
	}); err != nil {
		pgUow.RevertChanges()
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	updatedWorldDto, err := worldAppService.GetWorld(worldappsrv.GetWorldQuery{WorldId: worldIdDto})
	if err != nil {
		pgUow.RevertChanges()
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	pgUow.SaveChanges()
	c.JSON(http.StatusOK, updateWorldResponse(updatedWorldDto))
}
