package httpcontroller

import (
	"net/http"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/application/service/itemappservice"
	"github.com/gin-gonic/gin"
)

type httpPresenter struct {
	c *gin.Context
}

func NewHttpPresenter(c *gin.Context) itemappservice.Presenter {
	return httpPresenter{
		c: c,
	}
}

func (p httpPresenter) OnSuccess(jsonObj any) {
	p.c.JSON(http.StatusOK, jsonObj)
}
