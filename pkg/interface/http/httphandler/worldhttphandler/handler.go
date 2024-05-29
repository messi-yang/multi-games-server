package worldhttphandler

import (
	"net/http"

	"github.com/dum-dum-genius/zossi-server/pkg/application/usecase"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/httpsession"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/viewmodel"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type HttpHandler struct{}

func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}

func (httpHandler *HttpHandler) GetWorld(c *gin.Context) {
	worldIdDto, err := uuid.Parse(c.Param("worldId"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	pgUow := pguow.NewDummyUow()
	getWorldUseCase := usecase.ProvideGetWorldUseCase(pgUow)
	worldDto, err := getWorldUseCase.Execute(worldIdDto)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getWorldResponse(
		viewmodel.WorldViewModel(worldDto),
	))
}

func (httpHandler *HttpHandler) GetMyWorlds(c *gin.Context) {
	authorizedUserIdDto := httpsession.GetAuthorizedUserId(c)
	if authorizedUserIdDto == nil {
		c.String(http.StatusUnauthorized, "not authorized")
		return
	}

	pgUow := pguow.NewDummyUow()

	getMyWorldsUseCase := usecase.ProvideGetMyWorldsUseCase(pgUow)
	worldDtos, err := getMyWorldsUseCase.Execute(*authorizedUserIdDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	worldViewModels := lo.Map(worldDtos, func(worldDto dto.WorldDto, _ int) viewmodel.WorldViewModel {
		return viewmodel.WorldViewModel(worldDto)
	})

	c.JSON(http.StatusOK, getMyWorldsResponse(worldViewModels))
}

func (httpHandler *HttpHandler) CreateWorld(c *gin.Context) {
	authorizedUserIdDto := httpsession.GetAuthorizedUserId(c)
	if authorizedUserIdDto == nil {
		c.String(http.StatusUnauthorized, "not authorized")
		return
	}

	var requestBody createWorldRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	pgUow := pguow.NewUow()
	createWorldUseCase := usecase.ProvideCreateWorldUseCase(pgUow)
	worldDto, err := createWorldUseCase.Execute(*authorizedUserIdDto, requestBody.Name)
	if err != nil {
		pgUow.RevertChanges()
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	pgUow.SaveChanges()
	c.JSON(http.StatusOK, createWorldResponse(viewmodel.WorldViewModel(worldDto)))
}

func (httpHandler *HttpHandler) UpdateWorld(c *gin.Context) {
	authorizedUserIdDto := httpsession.GetAuthorizedUserId(c)
	if authorizedUserIdDto == nil {
		c.String(http.StatusUnauthorized, "not authorized")
		return
	}

	var requestBody updateWorldRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	worldIdDto, err := uuid.Parse(c.Param("worldId"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	pgUow := pguow.NewUow()
	updateWorldUseCase := usecase.ProvideUpdateWorldUseCase(pgUow)
	updatedWorldDto, err := updateWorldUseCase.Execute(*authorizedUserIdDto, worldIdDto, requestBody.Name)
	if err != nil {
		pgUow.RevertChanges()
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	pgUow.SaveChanges()

	c.JSON(http.StatusOK, updateWorldResponse(viewmodel.WorldViewModel(updatedWorldDto)))
}

func (httpHandler *HttpHandler) DeleteWorld(c *gin.Context) {
	authorizedUserIdDto := httpsession.GetAuthorizedUserId(c)
	if authorizedUserIdDto == nil {
		c.String(http.StatusUnauthorized, "not authorized")
		return
	}

	worldIdDto, err := uuid.Parse(c.Param("worldId"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	pgUow := pguow.NewUow()
	deleteWorldUseCase := usecase.ProvideDeleteWorldUseCase(pgUow)
	if err = deleteWorldUseCase.Execute(*authorizedUserIdDto, worldIdDto); err != nil {
		pgUow.RevertChanges()
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	pgUow.SaveChanges()

	c.String(http.StatusOK, "")
}
