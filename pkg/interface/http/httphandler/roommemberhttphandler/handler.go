package roommemberhttphandler

import (
	"net/http"

	"github.com/dum-dum-genius/zossi-server/pkg/application/usecase"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/httpsession"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HttpHandler struct{}

func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}

func (httpHandler *HttpHandler) GetRoomMembers(c *gin.Context) {
	authorizedUserIdDto := httpsession.GetAuthorizedUserId(c)
	if authorizedUserIdDto == nil {
		c.String(http.StatusUnauthorized, "not authorized")
		return
	}

	roomIdDto, err := uuid.Parse(c.Param("roomId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	pgUow := pguow.NewDummyUow()
	queryRoomMembersUseCase := usecase.ProvideQueryRoomMembersUseCase(pgUow)
	roomMemberDtos, err := queryRoomMembersUseCase.Execute(roomIdDto, *authorizedUserIdDto)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, roomMemberDtos)
}
