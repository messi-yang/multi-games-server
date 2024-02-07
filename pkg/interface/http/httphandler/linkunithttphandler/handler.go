package linkunithttphandler

import (
	"net/http"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/linkunitappsrv"
	world_provide_dependency "github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/providedependency"
	"github.com/gin-gonic/gin"
)

type HttpHandler struct{}

func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}

func (httpHandler *HttpHandler) GetLinkUnitUrl(c *gin.Context) {
	var requestBody getLinkUnitUrlRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	pgUow := pguow.NewDummyUow()

	linkUnitAppService := world_provide_dependency.ProvideLinkUnitAppService(pgUow)
	url, err := linkUnitAppService.GetLinkUnitUrl(linkunitappsrv.GetLinkUnitUrlQuery{
		WorldId:  requestBody.WorldId,
		Position: requestBody.Position,
	})
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getLinkUnitUrlResponse{
		Url: url,
	})
}
