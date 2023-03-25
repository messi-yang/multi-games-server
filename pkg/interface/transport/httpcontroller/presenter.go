package httpcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Presenter struct {
	ginContext *gin.Context
}

func NewPresenter(ginContext *gin.Context) *Presenter {
	return &Presenter{
		ginContext: ginContext,
	}
}

func (presenter Presenter) OnSuccess(jsonObj any) {
	presenter.ginContext.JSON(http.StatusOK, jsonObj)
}

func (presenter Presenter) OnError(err error) {
	presenter.ginContext.JSON(http.StatusBadRequest, err)
}
