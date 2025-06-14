package portalunithttphandler

import (
	"net/http"
	"strconv"

	"github.com/dum-dum-genius/zossi-server/pkg/application/usecase"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HttpHandler struct{}

func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}

func (httpHandler *HttpHandler) GetPortalUnitTargetUnit(c *gin.Context) {
	idDto, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	pgUow := pguow.NewDummyUow()
	getPortalUnitTargetUnitUseCase := usecase.ProvideGetPortalUnitTargetUnitUseCase(pgUow)
	targetPortalUnitDto, err := getPortalUnitTargetUnitUseCase.Execute(idDto)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, targetPortalUnitDto)
}

func (httpHandler *HttpHandler) QueryPortalUnits(c *gin.Context) {
	worldIdDto, err := uuid.Parse(c.Query("world_id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	limitString := c.Query("limit")
	if limitString == "" {
		limitString = "20"
	}
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	offsetString := c.Query("offset")
	if offsetString == "" {
		offsetString = "0"
	}
	offset, err := strconv.Atoi(offsetString)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	pgUow := pguow.NewDummyUow()
	queryPortalUnitsUseCase := usecase.ProvideQueryPortalUnitsUseCase(pgUow)
	portalUnitsDto, err := queryPortalUnitsUseCase.Execute(worldIdDto, limit, offset)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, portalUnitsDto)
}
