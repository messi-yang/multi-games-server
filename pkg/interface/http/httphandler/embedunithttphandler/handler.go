package embedunithttphandler

import (
	"net/http"

	"github.com/dum-dum-genius/zossi-server/pkg/application/usecase"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HttpHandler struct{}

func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}

func (httpHandler *HttpHandler) GetEmbedUnitEmbedCode(c *gin.Context) {
	idDto, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	pgUow := pguow.NewDummyUow()
	getEmbedUnitEmbedCodeUseCase := usecase.ProvideGetEmbedUnitEmbedCodeUseCase(pgUow)
	embedCodeDto, err := getEmbedUnitEmbedCodeUseCase.Execute(idDto)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getEmbedUnitResponse{
		EmbedCode: embedCodeDto,
	})
}
