package roomhttphandler

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

func (httpHandler *HttpHandler) GetRoom(c *gin.Context) {
	roomIdDto, err := uuid.Parse(c.Param("roomId"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	pgUow := pguow.NewDummyUow()
	getRoomUseCase := usecase.ProvideGetRoomUseCase(pgUow)
	roomDto, err := getRoomUseCase.Execute(roomIdDto)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getRoomResponse(
		viewmodel.RoomViewModel(roomDto),
	))
}

func (httpHandler *HttpHandler) GetMyRooms(c *gin.Context) {
	authorizedUserIdDto := httpsession.GetAuthorizedUserId(c)
	if authorizedUserIdDto == nil {
		c.String(http.StatusUnauthorized, "not authorized")
		return
	}

	pgUow := pguow.NewDummyUow()

	getMyRoomsUseCase := usecase.ProvideGetMyRoomsUseCase(pgUow)
	roomDtos, err := getMyRoomsUseCase.Execute(*authorizedUserIdDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	roomViewModels := lo.Map(roomDtos, func(roomDto dto.RoomDto, _ int) viewmodel.RoomViewModel {
		return viewmodel.RoomViewModel(roomDto)
	})

	c.JSON(http.StatusOK, getMyRoomsResponse(roomViewModels))
}

func (httpHandler *HttpHandler) CreateRoom(c *gin.Context) {
	authorizedUserIdDto := httpsession.GetAuthorizedUserId(c)
	if authorizedUserIdDto == nil {
		c.String(http.StatusUnauthorized, "not authorized")
		return
	}

	var requestBody createRoomRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	pgUow := pguow.NewUow()
	createRoomUseCase := usecase.ProvideCreateRoomUseCase(pgUow)
	roomDto, err := createRoomUseCase.Execute(*authorizedUserIdDto, requestBody.Name)
	if err != nil {
		pgUow.RevertChanges()
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	pgUow.SaveChanges()
	c.JSON(http.StatusOK, createRoomResponse(viewmodel.RoomViewModel(roomDto)))
}

func (httpHandler *HttpHandler) UpdateRoom(c *gin.Context) {
	authorizedUserIdDto := httpsession.GetAuthorizedUserId(c)
	if authorizedUserIdDto == nil {
		c.String(http.StatusUnauthorized, "not authorized")
		return
	}

	var requestBody updateRoomRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	roomIdDto, err := uuid.Parse(c.Param("roomId"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	pgUow := pguow.NewUow()
	updateRoomUseCase := usecase.ProvideUpdateRoomUseCase(pgUow)
	updatedRoomDto, err := updateRoomUseCase.Execute(*authorizedUserIdDto, roomIdDto, requestBody.Name)
	if err != nil {
		pgUow.RevertChanges()
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	pgUow.SaveChanges()

	c.JSON(http.StatusOK, updateRoomResponse(viewmodel.RoomViewModel(updatedRoomDto)))
}

func (httpHandler *HttpHandler) DeleteRoom(c *gin.Context) {
	authorizedUserIdDto := httpsession.GetAuthorizedUserId(c)
	if authorizedUserIdDto == nil {
		c.String(http.StatusUnauthorized, "not authorized")
		return
	}

	roomIdDto, err := uuid.Parse(c.Param("roomId"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	pgUow := pguow.NewUow()
	deleteRoomUseCase := usecase.ProvideDeleteRoomUseCase(pgUow)
	if err = deleteRoomUseCase.Execute(*authorizedUserIdDto, roomIdDto); err != nil {
		pgUow.RevertChanges()
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	pgUow.SaveChanges()

	c.String(http.StatusOK, "")
}
