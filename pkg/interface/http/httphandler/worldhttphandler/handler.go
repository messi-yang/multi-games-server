package worldhttphandler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/application/service/userappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/application/service/worldmemberappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/application/service/worldpermissionappsrv"
	iam_provide_dependency "github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/providedependency"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/worldappsrv"
	world_provide_dependency "github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/providedependency"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/httputil"
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
	userAppService := iam_provide_dependency.ProvideUserAppService(pgUow)

	worldDto, err := worldAppService.GetWorld(worldappsrv.GetWorldQuery{WorldId: worldIdDto})
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	userDto, err := userAppService.GetUser(userappsrv.GetUserQuery{UserId: worldDto.UserId})
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getWorldResponse(
		viewmodel.WorldViewModel{
			WorldDto: worldDto,
			UserDto:  userDto,
		},
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
	userAppService := iam_provide_dependency.ProvideUserAppService(pgUow)

	worldDtos, err := worldAppService.QueryWorlds(worldappsrv.QueryWorldsQuery{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	userIdDtos := lo.Map(worldDtos, func(worldDto dto.WorldDto, _ int) uuid.UUID {
		return worldDto.UserId
	})
	userDtoMap, err := userAppService.GetUsersOfIds(userappsrv.GetUsersOfIdsQuery{UserIds: userIdDtos})
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	worldViewModels := lo.Map(worldDtos, func(worldDto dto.WorldDto, _ int) viewmodel.WorldViewModel {
		return viewmodel.WorldViewModel{
			WorldDto: worldDto,
			UserDto:  userDtoMap[worldDto.UserId],
		}
	})

	c.JSON(http.StatusOK, queryWorldsResponse(worldViewModels))
}

func (httpHandler *HttpHandler) GetMyWorlds(c *gin.Context) {
	userIdDto := httputil.GetUserId(c)

	pgUow := pguow.NewDummyUow()

	worldAppService := world_provide_dependency.ProvideWorldAppService(pgUow)
	userAppService := iam_provide_dependency.ProvideUserAppService(pgUow)

	worldDtos, err := worldAppService.GetMyWorlds(worldappsrv.GetMyWorldsQuery{
		UserId: userIdDto,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userIdDtos := lo.Map(worldDtos, func(worldDto dto.WorldDto, _ int) uuid.UUID {
		return worldDto.UserId
	})
	userDtoMap, err := userAppService.GetUsersOfIds(userappsrv.GetUsersOfIdsQuery{UserIds: userIdDtos})
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	worldViewModels := lo.Map(worldDtos, func(worldDto dto.WorldDto, _ int) viewmodel.WorldViewModel {
		return viewmodel.WorldViewModel{
			WorldDto: worldDto,
			UserDto:  userDtoMap[worldDto.UserId],
		}
	})

	c.JSON(http.StatusOK, getMyWorldsResponse(worldViewModels))
}

func (httpHandler *HttpHandler) CreateWorld(c *gin.Context) {
	userIdDto := httputil.GetUserId(c)

	var requestBody createWorldRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	pgUow := pguow.NewUow()

	worldAppService := world_provide_dependency.ProvideWorldAppService(pgUow)
	worldMemberAppService := iam_provide_dependency.ProvideWorldMemberAppService(pgUow)
	userAppService := iam_provide_dependency.ProvideUserAppService(pgUow)

	newWorldIdDto, err := worldAppService.CreateWorld(
		worldappsrv.CreateWorldCommand{
			UserId: userIdDto,
			Name:   requestBody.Name,
		},
	)
	if err != nil {
		pgUow.RevertChanges()
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// TODO - handle this side effects by using integration events
	if err := worldMemberAppService.AddWorldMember(worldmemberappsrv.AddWorldMemberCommand{
		UserId:  userIdDto,
		WorldId: newWorldIdDto,
		Role:    "owner",
	}); err != nil {
		pgUow.RevertChanges()
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	worldDto, err := worldAppService.GetWorld(worldappsrv.GetWorldQuery{WorldId: newWorldIdDto})
	if err != nil {
		pgUow.RevertChanges()
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	userDto, err := userAppService.GetUser(userappsrv.GetUserQuery{UserId: worldDto.UserId})
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	pgUow.SaveChanges()
	c.JSON(http.StatusOK, createWorldResponse(viewmodel.WorldViewModel{
		WorldDto: worldDto,
		UserDto:  userDto,
	}))
}

func (httpHandler *HttpHandler) UpdateWorld(c *gin.Context) {
	userIdDto := httputil.GetUserId(c)

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
	userAppService := iam_provide_dependency.ProvideUserAppService(pgUow)

	canUpdateWorld, err := worldPermissionAppService.CanUpdateWorld(worldpermissionappsrv.CanUpdateWorldQuery{
		WorldId: worldIdDto,
		UserId:  userIdDto,
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

	userDto, err := userAppService.GetUser(userappsrv.GetUserQuery{UserId: updatedWorldDto.UserId})
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	pgUow.SaveChanges()
	c.JSON(http.StatusOK, updateWorldResponse(viewmodel.WorldViewModel{
		WorldDto: updatedWorldDto,
		UserDto:  userDto,
	}))
}

func (httpHandler *HttpHandler) DeleteWorld(c *gin.Context) {
	userIdDto := httputil.GetUserId(c)

	worldIdDto, err := uuid.Parse(c.Param("worldId"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	pgUow := pguow.NewUow()

	worldAppService := world_provide_dependency.ProvideWorldAppService(pgUow)
	worldMemberAppService := iam_provide_dependency.ProvideWorldMemberAppService(pgUow)
	worldPermissionAppService := iam_provide_dependency.ProvideWorldPermissionAppService(pgUow)

	canDeleteWorld, err := worldPermissionAppService.CanDeleteWorld(worldpermissionappsrv.CanDeleteWorldQuery{
		WorldId: worldIdDto,
		UserId:  userIdDto,
	})
	if err != nil {
		pgUow.RevertChanges()
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if !canDeleteWorld {
		pgUow.RevertChanges()
		c.String(http.StatusForbidden, "not permitted")
		return
	}

	// TODO - handle this side effects by using integration events
	if err := worldMemberAppService.DeleteAllWorldMembersInWorld(worldmemberappsrv.DeleteAllWorldMembersInWorldCommand{
		WorldId: worldIdDto,
	}); err != nil {
		pgUow.RevertChanges()
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err = worldAppService.DeleteWorld(worldappsrv.DeleteWorldCommand{
		WorldId: worldIdDto,
	}); err != nil {
		pgUow.RevertChanges()
		if errors.Is(err, worldappsrv.ErrNotPermitted) {
			c.String(http.StatusForbidden, err.Error())
		} else {
			c.String(http.StatusBadRequest, err.Error())
		}
		return
	}

	pgUow.SaveChanges()
	c.String(http.StatusOK, "")
}
