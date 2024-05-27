package worldhttphandler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/dum-dum-genius/zossi-server/pkg/application/usecase"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/application/service/worldpermissionappsrv"
	iam_provide_dependency "github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/providedependency"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/worldappsrv"
	world_provide_dependency "github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/providedependency"
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

	worldAppService := world_provide_dependency.ProvideWorldAppService(pgUow)

	worldDto, err := worldAppService.GetWorld(worldappsrv.GetWorldQuery{WorldId: worldIdDto})
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getWorldResponse(
		viewmodel.WorldViewModel(worldDto),
	))
}

func (httpHandler *HttpHandler) QueryWorlds(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	offsetQuery := c.Query("offset")
	offset, err := strconv.Atoi(offsetQuery)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	pgUow := pguow.NewDummyUow()

	worldAppService := world_provide_dependency.ProvideWorldAppService(pgUow)

	worldDtos, err := worldAppService.QueryWorlds(worldappsrv.QueryWorldsQuery{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	worldViewModels := lo.Map(worldDtos, func(worldDto dto.WorldDto, _ int) viewmodel.WorldViewModel {
		return viewmodel.WorldViewModel(worldDto)
	})

	c.JSON(http.StatusOK, queryWorldsResponse(worldViewModels))
}

func (httpHandler *HttpHandler) GetMyWorlds(c *gin.Context) {
	authorizedUserIdDto := httpsession.GetAuthorizedUserId(c)
	if authorizedUserIdDto == nil {
		c.String(http.StatusUnauthorized, "not authorized")
		return
	}

	pgUow := pguow.NewDummyUow()

	worldAppService := world_provide_dependency.ProvideWorldAppService(pgUow)

	worldDtos, err := worldAppService.GetMyWorlds(worldappsrv.GetMyWorldsQuery{
		UserId: *authorizedUserIdDto,
	})
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

	worldAppService := world_provide_dependency.ProvideWorldAppService(pgUow)
	worldPermissionAppService := iam_provide_dependency.ProvideWorldPermissionAppService(pgUow)

	canUpdateWorld, err := worldPermissionAppService.CanUpdateWorld(worldpermissionappsrv.CanUpdateWorldQuery{
		WorldId: worldIdDto,
		UserId:  *authorizedUserIdDto,
	})
	if err != nil {
		pgUow.RevertChanges()
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if !canUpdateWorld {
		pgUow.RevertChanges()
		c.String(http.StatusForbidden, "not permitted")
		return
	}

	if err = worldAppService.UpdateWorld(worldappsrv.UpdateWorldCommand{
		WorldId: worldIdDto,
		Name:    requestBody.Name,
	}); err != nil {
		pgUow.RevertChanges()
		if errors.Is(err, worldappsrv.ErrNotPermitted) {
			c.String(http.StatusForbidden, err.Error())
		} else {
			c.String(http.StatusBadRequest, err.Error())
		}
		return
	}

	updatedWorldDto, err := worldAppService.GetWorld(worldappsrv.GetWorldQuery{WorldId: worldIdDto})
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
