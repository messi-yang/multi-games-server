package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpPresenter struct {
	ginContext *gin.Context
}

func NewHttpPresenter(ginContext *gin.Context) *HttpPresenter {
	return &HttpPresenter{
		ginContext: ginContext,
	}
}

func (presenter HttpPresenter) OnSuccess(jsonObj any) {
	presenter.ginContext.JSON(http.StatusOK, jsonObj)
}

func (presenter HttpPresenter) OnError(err error) {
	presenter.ginContext.JSON(http.StatusBadRequest, err.Error())
}
