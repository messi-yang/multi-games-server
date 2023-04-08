package userhttpcontroller

import (
	"net/http"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/userappservice"
	"github.com/gin-gonic/gin"
)

func getUsersHandler(c *gin.Context) {
	userAppService, err := provideUserAppService()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userDtos, err := userAppService.GetUsers(userappservice.GetUsersQuery{})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getUsersReponseDto(userDtos))
}
