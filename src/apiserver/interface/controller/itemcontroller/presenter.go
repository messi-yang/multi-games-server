package itemcontroller

import (
	"net/http"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/application/service/itemappservice"
	"github.com/gin-gonic/gin"
)

type ginPresenter struct {
	c *gin.Context
}

func newGinPresenter(c *gin.Context) itemappservice.Presenter {
	return ginPresenter{
		c: c,
	}
}

func (p ginPresenter) OnSuccess(jsonObj any) {
	p.c.JSON(http.StatusOK, jsonObj)
}
