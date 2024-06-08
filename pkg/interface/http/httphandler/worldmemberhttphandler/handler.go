package worldmemberhttphandler

import (
	"net/http"

	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/application/usecase"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
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

func (httpHandler *HttpHandler) GetWorldMembers(c *gin.Context) {
	authorizedUserIdDto := httpsession.GetAuthorizedUserId(c)
	if authorizedUserIdDto == nil {
		c.String(http.StatusUnauthorized, "not authorized")
		return
	}

	worldIdDto, err := uuid.Parse(c.Param("worldId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	pgUow := pguow.NewDummyUow()
	queryWorldMembersUseCase := usecase.ProvideQueryWorldMembersUseCase(pgUow)
	worldMemberDtos, err := queryWorldMembersUseCase.Execute(worldIdDto, *authorizedUserIdDto)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	worldMemberViewModels := lo.Map(worldMemberDtos, func(worldMemberDto dto.WorldMemberDto, _ int) viewmodel.WorldMemberViewModel {
		return viewmodel.WorldMemberViewModel(worldMemberDto)
	})

	c.JSON(http.StatusOK, getWorldMembersResponse(worldMemberViewModels))
}
